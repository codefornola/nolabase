package norta

import (
	"io"
	"net/http"
	"os"

	"github.com/codefornola/nolabase/internal/infra"
	"github.com/geops/gtfsparser"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

const (
	ScraperName = "NORTA"
	NortaURL    = "https://simst.im/static/norta-gtfs.zip"
)

var log = logrus.WithField("package", "norta")

type NortaScraper struct {
	repo *Repo
}

func NewScraper() *NortaScraper {
	return &NortaScraper{}
}

func (s *NortaScraper) Configure(pool *pgxpool.Pool) (err error) {
	s.repo = NewRepo(pool)
	return err
}

func (s *NortaScraper) EnqueueJobs() ([]infra.Job, error) {
	var jobs []infra.Job

	job := infra.Job{
		Url:         NortaURL,
		ScraperName: ScraperName,
	}
	jobs = append(jobs, job)

	return jobs, nil
}

func (s *NortaScraper) MakeRequest(j infra.Job) (*http.Request, error) {
	req, err := http.NewRequest("GET", j.Url, nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (s *NortaScraper) HandleResponse(j infra.Job, resp *http.Response, httpErr error) (*infra.Job, error) {
	scraperLog := log.WithField("job", j.Id).WithField("url", resp.Request.URL.String())
	// if we have an http error just return it
	if httpErr != nil {
		return nil, httpErr
	}

	// I don't think there is a way to parse bytes
	// directly so let's write to file first.
	// TODO fix
	defer resp.Body.Close()
	out, err := os.Create("/tmp/norta-gtfs.zip")
	if err != nil {
		return nil, err
	}
	defer out.Close()
	io.Copy(out, resp.Body)

	feed := gtfsparser.NewFeed()
	err = feed.Parse("/tmp/norta-gtfs.zip")
	if err != nil {
		return nil, err
	}

	err = s.repo.StoreGtfs(feed)
	if err != nil {
		scraperLog.Error(err)
		return nil, err
	}

	return nil, err
}
