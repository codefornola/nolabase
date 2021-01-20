package scraper

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/codefornola/nolabase/internal/datasources/assessor"
	"github.com/codefornola/nolabase/internal/datasources/callsforservice"
	"github.com/codefornola/nolabase/internal/datasources/councildistricts"
	"github.com/codefornola/nolabase/internal/datasources/neighborhoods"
	"github.com/codefornola/nolabase/internal/datasources/norta"
	"github.com/codefornola/nolabase/internal/datasources/policedistricts"
	"github.com/codefornola/nolabase/internal/datasources/policesubzones"
	"github.com/codefornola/nolabase/internal/datasources/restaurants"
	"github.com/codefornola/nolabase/internal/datasources/schooldistricts"
	"github.com/codefornola/nolabase/internal/datasources/shorttermrentals"
	"github.com/codefornola/nolabase/internal/datasources/vacationrentals"
	"github.com/codefornola/nolabase/internal/datasources/votingprecincts"
	"github.com/codefornola/nolabase/internal/infra"
	"github.com/gammazero/workerpool"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

const jobPgChannel = "new_jobs"

type JobNotificationPayload struct {
	Data struct {
		Id int
	}
}

var log = logrus.WithField("package", "scraper")

type ScraperBuilder func() Scraper

var AllScrapers = map[string]ScraperBuilder{
	(assessor.IdScraperName):       func() Scraper { return assessor.NewIdScraper() },
	(assessor.PropertyScraperName): func() Scraper { return assessor.NewPropertyScraper() },
	(callsforservice.ScraperName):  func() Scraper { return callsforservice.NewScraper() },
	(vacationrentals.ScraperName):  func() Scraper { return vacationrentals.NewScraper() },
	(shorttermrentals.ScraperName): func() Scraper { return shorttermrentals.NewScraper() },
	(neighborhoods.ScraperName):    func() Scraper { return neighborhoods.NewScraper() },
	(councildistricts.ScraperName): func() Scraper { return councildistricts.NewScraper() },
	(votingprecincts.ScraperName):  func() Scraper { return votingprecincts.NewScraper() },
	(schooldistricts.ScraperName):  func() Scraper { return schooldistricts.NewScraper() },
	(policedistricts.ScraperName):  func() Scraper { return policedistricts.NewScraper() },
	(policesubzones.ScraperName):   func() Scraper { return policesubzones.NewScraper() },
	(restaurants.ScraperName):      func() Scraper { return restaurants.NewScraper() },
	(norta.ScraperName):            func() Scraper { return norta.NewScraper() },
}

type HttpCall func() (*http.Response, error)

type Scraper interface {
	Configure(pool *pgxpool.Pool) error
	EnqueueJobs() ([]infra.Job, error)
	MakeRequest(infra.Job) (*http.Request, error)
	HandleResponse(infra.Job, *http.Response, error) (*infra.Job, error)
}

type ScraperManager struct {
	repo       *ScraperRepo
	httpClient *http.Client
	jobRepo    *infra.JobRepo
}

func NewScraperManager(dbUrl string) (*ScraperManager, error) {
	pool, err := pgxpool.Connect(context.Background(), dbUrl)
	if err != nil {
		return nil, err
	}
	repo := NewScraperRepo(pool)
	client := &http.Client{
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout: 30 * time.Second,
				// longer keepalive than default 10 seconds
				KeepAlive: 10 * time.Minute,
				DualStack: true,
			}).DialContext,
		},
	}
	jobRepo := infra.NewJobRepo(pool)
	return &ScraperManager{
		httpClient: client,
		repo:       repo,
		jobRepo:    jobRepo,
	}, nil
}

func (s *ScraperManager) GetConnectionPool() *pgxpool.Pool {
	return s.repo.pgConn
}

func (s *ScraperManager) SetTransport(transport *http.Transport) {
	s.httpClient.Transport = transport
}

func (s *ScraperManager) GetScraper(name string) (Scraper, error) {
	name = strings.ToUpper(name)
	if builder, ok := AllScrapers[name]; ok {
		return builder(), nil
	}
	msg := fmt.Sprintf("Couldn't find scraper with name %s", name)
	return nil, errors.New(msg)
}

func (s *ScraperManager) StartDatabaseEnqueuer(jobs chan infra.Job) {
	conn, err := s.repo.pgConn.Acquire(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Release()

	// TODO need to rethink this strategy as we could miss some stuff

	log.Info("Looking for existing jobs that are ready to be processed")
	// first query for existing jobs
	s.repo.ListReadyJobs(func(j *infra.Job) error {
		jobs <- *j
		return nil
	})
	log.Info("Done looking for existing jobs")

	_, err = conn.Exec(context.Background(), "listen "+jobPgChannel)
	if err != nil {
		log.Error("Failure listening")
		log.Fatal(err)
	}

	log.Info("Listening for new jobs...")
	for {
		notification, err := conn.Conn().WaitForNotification(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		var payload JobNotificationPayload
		err = json.Unmarshal([]byte(notification.Payload), &payload)
		if err != nil {
			log.Println("Failure unmarshalling")
			log.Fatal(err)
		}
		log.Debug("Got notification: ", payload)
		job, err := s.jobRepo.GetJob(payload.Data.Id)
		if err != nil {
			log.Fatal(err)
		}
		jobs <- job
	}
}

func (s *ScraperManager) StartSchedulerDaemon(concurrency int, jobs chan infra.Job) {
	scrapePool := workerpool.New(concurrency)
	for j := range jobs {
		if j.Id == 0 {
			log.Fatal("Scheduler got a non-persisted job ", j)
		}
		jobLog := log.WithField("job", j.Id)
		jobLog.Debug("Marking job as queued")
		err := s.repo.MarkJob(j.Id, infra.JobQueued)
		jobLog.Debug("Marked job as queued")
		if err != nil {
			jobLog.Error(err)
		}
		jobLog.Debug("Adding started job to the pool")
		scrapePool.Submit(s.newWorker(j))
	}
}

func (s *ScraperManager) newWorker(j infra.Job) func() {
	return func() {
		jobLog := log.WithField("job", j.Id)
		newJob, err := s.RunJob(j)
		if err != nil {
			jobLog.Debug("Marking job as failed because: ", err)
			markErr := s.repo.MarkJobFailed(j.Id, fmt.Sprintf("%+v", err))
			jobLog.Debug("Marked job as failed")
			if markErr != nil {
				jobLog.Error("failed mark job as faield, moving on anyway")
			}
			return
		}
		if newJob != nil {
			jobLog.Debug("Adding an extra job")
			_, err := s.repo.PersistJobs([]infra.Job{*newJob})
			if err != nil {
				jobLog.Error("Could not queue job ", newJob)
				return
			}
			jobLog.Debug("Added extra job")
		}
		jobLog.Debug("Marking job succeeded")
		err = s.repo.MarkJob(j.Id, infra.JobSucceeded)
		if err != nil {
			jobLog.Error("Error marking job succeeded: ", err)
			return
		}
	}
}

func (s *ScraperManager) Enqueue(scrapers []string) {
	for _, name := range scrapers {
		err := s.EnqueueJobs(name)
		if err != nil {
			log.Error("Error with scraper ", err)
		}
	}
}

func (s *ScraperManager) EnqueueJobs(scraperName string) error {
	scraperLog := log.WithField("scraper", scraperName)
	scraper, err := s.GetScraper(scraperName)
	if err != nil {
		return err
	}
	conn := s.repo.pgConn
	err = scraper.Configure(conn)
	if err != nil {
		return err
	}
	scraperLog.Debug("Preparing...")
	newJobs, err := scraper.EnqueueJobs()
	if err != nil {
		return err
	}

	scraperLog.Debug("Asked by scraper to enqueue jobs: ", len(newJobs))
	_, err = s.repo.PersistJobs(newJobs)
	if err != nil {
		return err
	}

	return err
}

func (s *ScraperManager) RunJob(j infra.Job) (*infra.Job, error) {
	jobLog := log.WithField("job", j.Id)
	err := s.repo.MarkJob(j.Id, infra.JobStarted)
	if err != nil {
		jobLog.Error("Couldn't mark job as started but moving on anyway")
	}
	scraper, err := s.GetScraper(j.ScraperName)
	if err != nil {
		return nil, err
	}
	err = scraper.Configure(s.repo.pgConn)
	if err != nil {
		return nil, err
	}
	jobLog.Debug("Making the request object")
	req, err := scraper.MakeRequest(j)
	if err != nil {
		return nil, err
	}
	jobLog.Debug("Executing Request:", req.URL.String())
	start := time.Now()
	resp, err := s.retry(5, 3*time.Second, func() (*http.Response, error) {
		return s.httpClient.Do(req)
	})
	jobLog.Debug("HTTP time ms: ", time.Since(start).Milliseconds())
	jobLog.Debug("Calling Handle Response")
	return scraper.HandleResponse(j, resp, err)
}

func (s *ScraperManager) retry(attempts int, sleep time.Duration, call HttpCall) (r *http.Response, err error) {
	for i := 0; ; i++ {
		result, err := call()
		if err == nil {
			return result, nil
		}

		if i >= (attempts - 1) {
			break
		}

		log.Printf("Sleeping %dms until retrying request\n", sleep.Milliseconds())
		time.Sleep(sleep)
		log.Println("Retrying request after error:", err)
	}
	return nil, fmt.Errorf("after %d attempts, last error: %s", attempts, err)
}
