package assessor

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/bhelx/nolabase/infra"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	searchUrl       = "http://qpublic9.qpublic.net/la_orleans_alsearch.php"
	parcelColumnNum = 1
	IdScraperName   = "ASSESSOR-IDS"
)

type JobMetadata struct {
	Term  string
	Begin int
}

type IdScraper struct {
	repo *Repo
}

func NewIdScraper() *IdScraper {
	return &IdScraper{}
}

func (s *IdScraper) Configure(pool *pgxpool.Pool) (err error) {
	s.repo = NewRepo(pool)
	return err
}

func (s *IdScraper) EnqueueJobs() ([]infra.Job, error) {
	var jobs []infra.Job

	meta := &JobMetadata{
		Term:  "E",
		Begin: 0,
	}
	metas, err := json.Marshal(&meta)
	if err != nil {
		return jobs, err
	}
	job := infra.Job{
		Url:         searchUrl,
		MetaData:    string(metas),
		ScraperName: IdScraperName,
	}
	jobs = append(jobs, job)

	// 0 - 9
	for i := 48; i <= 57; i++ {
		meta := &JobMetadata{
			Term:  string(rune(i)),
			Begin: 0,
		}
		metas, err := json.Marshal(&meta)
		if err != nil {
			return jobs, err
		}
		job := infra.Job{
			Url:         searchUrl,
			MetaData:    string(metas),
			ScraperName: IdScraperName,
		}
		jobs = append(jobs, job)
	}

	// A - Z
	for i := 65; i <= 90; i++ {
		meta := &JobMetadata{
			Term:  string(rune(i)),
			Begin: 0,
		}
		metas, err := json.Marshal(&meta)
		if err != nil {
			return jobs, err
		}
		job := infra.Job{
			Url:         searchUrl,
			MetaData:    string(metas),
			ScraperName: IdScraperName,
		}
		jobs = append(jobs, job)
	}

	return jobs, nil
}

func (s *IdScraper) MakeRequest(j infra.Job) (*http.Request, error) {
	meta := &JobMetadata{}
	err := json.Unmarshal([]byte(j.MetaData), meta)
	if err != nil {
		return nil, err
	}
	body := fmt.Sprintf("INPUT=%s&BEGIN=%d", meta.Term, meta.Begin)
	body += "&searchType=owner_name&Owner_Search=Search%20By%20Owner%20Name"
	req, err := http.NewRequest("POST", searchUrl, bytes.NewReader([]byte(body)))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return req, nil
}

func (s *IdScraper) HandleResponse(j infra.Job, resp *http.Response, httpErr error) (*infra.Job, error) {
	if httpErr != nil {
		return nil, httpErr
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	var properties []*Property
	doc.Find("table > tbody > tr").Each(func(i int, s *goquery.Selection) {
		parcel := s.Find(fmt.Sprintf("td.search_value:nth-child(%d)", parcelColumnNum)).Text()
		parcel = trimTaxBill(parcel)
		if parcel != "" {
			properties = append(properties, &Property{AssessorId: parcel})
		}
	})

	if len(properties) > 0 {
		err := s.repo.StoreNewProperties(properties)
		if err != nil {
			return nil, err
		}
	}

	// if we have 100, we should try to fetch a new page
	if len(properties) == 100 {
		meta := &JobMetadata{}
		err = json.Unmarshal([]byte(j.MetaData), meta)
		if err != nil {
			return nil, err
		}
		newMeta := &JobMetadata{
			Term:  meta.Term,
			Begin: meta.Begin + 100,
		}
		metas, err := json.Marshal(newMeta)
		if err != nil {
			return nil, err
		}
		job := &infra.Job{
			Url:         searchUrl,
			MetaData:    string(metas),
			ScraperName: IdScraperName,
		}
		return job, nil
	}

	return nil, nil
}

func trimTaxBill(s string) string {
	return strings.TrimSpace(s)
}
