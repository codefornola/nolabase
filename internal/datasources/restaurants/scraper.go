package restaurants

import (
	"bytes"
	"compress/gzip"
	"io"
	"net/http"

	"github.com/codefornola/nolabase/internal/infra"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

const (
	ScraperName    = "RESTAURANTS"
	RestaurantsURL = "https://opendata.arcgis.com/datasets/62501c695c614d1c99aaf2b113fca257_0.geojson"
)

var log = logrus.WithField("package", "restaurants")

type RestaurantsScraper struct {
	repo *Repo
}

func NewScraper() *RestaurantsScraper {
	return &RestaurantsScraper{}
}

func (s *RestaurantsScraper) Configure(pool *pgxpool.Pool) (err error) {
	s.repo = NewRepo(pool)
	return err
}

func (s *RestaurantsScraper) EnqueueJobs() (jobs []infra.Job, err error) {
	job := infra.Job{
		Url:         RestaurantsURL,
		ScraperName: ScraperName,
	}
	jobs = append(jobs, job)
	return jobs, nil
}

func (s *RestaurantsScraper) MakeRequest(j infra.Job) (*http.Request, error) {
	req, err := http.NewRequest("GET", j.Url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept-Encoding", "deflate, gzip;q=1.0, *;q=0.5")
	return req, nil
}

func (s *RestaurantsScraper) HandleResponse(j infra.Job, resp *http.Response, httpErr error) (*infra.Job, error) {
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

	restaurants, err := ParseRestaurants(body.Bytes())
	if err != nil {
		return nil, err
	}

	err = s.repo.StoreRestaurantss(restaurants)
	return nil, err
}
