package policedistricts

import (
	"bytes"
	"compress/gzip"
	"io"
	"net/http"

	"github.com/bhelx/nolabase/infra"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

const (
	ScraperName        = "POLICE-DISTRICTS"
	PoliceDistrictsURL = "https://opendata.arcgis.com/datasets/140759858aa14bb6a5a2fe099ccf4c07_0.geojson"
)

var log = logrus.WithField("package", "policedistricts")

type PoliceDistrictsScraper struct {
	repo *Repo
}

func NewScraper() *PoliceDistrictsScraper {
	return &PoliceDistrictsScraper{}
}

func (s *PoliceDistrictsScraper) Configure(pool *pgxpool.Pool) (err error) {
	s.repo = NewRepo(pool)
	return err
}

func (s *PoliceDistrictsScraper) EnqueueJobs() ([]infra.Job, error) {
	var jobs []infra.Job

	job := infra.Job{
		Url:         PoliceDistrictsURL,
		ScraperName: ScraperName,
	}
	jobs = append(jobs, job)

	return jobs, nil
}

func (s *PoliceDistrictsScraper) MakeRequest(j infra.Job) (*http.Request, error) {
	req, err := http.NewRequest("GET", j.Url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept-Encoding", "deflate, gzip;q=1.0, *;q=0.5")
	return req, nil
}

func (s *PoliceDistrictsScraper) HandleResponse(j infra.Job, resp *http.Response, httpErr error) (*infra.Job, error) {
	scraperLog := log.WithField("job", j.Id).WithField("url", resp.Request.URL.String())
	// if we have an http error just return it
	if httpErr != nil {
		return nil, httpErr
	}

	var err error
	var reader io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(resp.Body)
		if err != nil {
			scraperLog.Error(err)
			return nil, err
		}
	default:
		reader = resp.Body
	}
	defer reader.Close()

	body := bytes.NewBuffer(nil)
	_, err = io.Copy(body, reader)
	if err != nil {
		scraperLog.Error(err)
		return nil, err
	}

	districts, err := ParsePoliceDistricts(body.Bytes())
	if err != nil {
		scraperLog.Error(err)
		return nil, err
	}

	err = s.repo.StorePoliceDistricts(districts)
	if err != nil {
		scraperLog.Error(err)
		return nil, err
	}

	return nil, err
}
