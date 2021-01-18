package assessor

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
	"github.com/codefornola/nolabase/internal/infra"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	parcelUrl           = "http://qpublic9.qpublic.net/la_orleans_display.php"
	PropertyScraperName = "ASSESSOR-PROPERTIES"
)

type PropertyPage struct {
	property *Property
	values   []*PropertyValue
	sales    []*PropertySale
}

type PropertyScraper struct {
	repo *Repo
}

func NewPropertyScraper() *PropertyScraper {
	return &PropertyScraper{}
}

func (s *PropertyScraper) Configure(pool *pgxpool.Pool) (err error) {
	s.repo = NewRepo(pool)
	return err
}

func (s *PropertyScraper) EnqueueJobs() ([]infra.Job, error) {
	var jobs []infra.Job
	s.repo.AllAssesorIds(func(id string) error {
		url := fmt.Sprintf("%s?KEY=%s", parcelUrl, id)
		job := infra.Job{
			Url:         url,
			ScraperName: PropertyScraperName,
		}
		jobs = append(jobs, job)
		return nil
	})

	return jobs, nil
}

func (s *PropertyScraper) MakeRequest(j infra.Job) (*http.Request, error) {
	req, err := http.NewRequest("GET", j.Url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept-Encoding", "deflate, gzip;q=1.0, *;q=0.5")
	return req, nil
}

func (s *PropertyScraper) HandleResponse(_ infra.Job, resp *http.Response, err error) (*infra.Job, error) {
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	url, err := url.Parse(resp.Request.URL.String())
	if err != nil {
		return nil, err
	}
	key := url.Query().Get("KEY")
	if key == "" {
		return nil, errors.New("Could not find assessor id in Request URL")
	}

	property := ParseProperty(doc)
	property.AssessorId = key

	page := &PropertyPage{
		property: property,
		values:   ParseValues(doc),
		sales:    ParseSales(doc),
	}

	return nil, s.repo.StorePropertyPage(page)
}
