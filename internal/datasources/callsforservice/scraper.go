package callsforservice

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/codefornola/nolabase/internal/infra"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

const (
	ScraperName  = "CALLS-FOR-SERVICE"
	CFS2011URL   = "https://data.nola.gov/resource/28ec-c8d6.json"
	CFS2012URL   = "https://data.nola.gov/resource/rv3g-ypg7.json"
	CFS2013URL   = "https://data.nola.gov/resource/5fn8-vtui.json"
	CFS2014URL   = "https://data.nola.gov/resource/jsyu-nz5r.json"
	CFS2015URL   = "https://data.nola.gov/resource/w68y-xmk6.json"
	CFS2016URL   = "https://data.nola.gov/resource/wgrp-d3ma.json"
	CFS2017URL   = "https://data.nola.gov/resource/bqmt-f3jk.json"
	CFS2018URL   = "https://data.nola.gov/resource/9san-ivhk.json"
	CFS2019URL   = "https://data.nola.gov/resource/qf6q-pp4b.json"
	CFS2020URL   = "https://data.nola.gov/resource/hp7u-i9hf.json"
	LimitPerPage = 1000
)

var log = logrus.WithField("package", "callsforservice")

type JobMetadata struct {
	Limit  int
	Offset int
}

type CfsScraper struct {
	repo *Repo
}

func NewScraper() *CfsScraper {
	return &CfsScraper{}
}

func (s *CfsScraper) Configure(pool *pgxpool.Pool) (err error) {
	s.repo = NewRepo(pool)
	return err
}

func (s *CfsScraper) newJob(url string) (j infra.Job, err error) {
	metab, err := json.Marshal(&JobMetadata{
		Limit:  LimitPerPage,
		Offset: 0,
	})
	if err != nil {
		return j, err
	}
	return infra.Job{
		Url:         url,
		ScraperName: ScraperName,
		MetaData:    string(metab),
	}, nil
}

func (s *CfsScraper) EnqueueJobs() ([]infra.Job, error) {
	var jobs []infra.Job

	urls := []string{
		CFS2011URL,
		CFS2012URL,
		CFS2013URL,
		CFS2014URL,
		CFS2015URL,
		CFS2016URL,
		CFS2017URL,
		CFS2018URL,
		CFS2019URL,
		CFS2020URL,
	}
	for _, u := range urls {
		job, err := s.newJob(u)
		if err != nil {
			return jobs, err
		}
		jobs = append(jobs, job)
	}

	return jobs, nil
}

func (s *CfsScraper) MakeRequest(j infra.Job) (*http.Request, error) {
	var meta JobMetadata
	err := json.Unmarshal([]byte(j.MetaData), &meta)
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s?$limit=%d&$offset=%d", j.Url, meta.Limit, meta.Offset)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept-Encoding", "deflate, gzip;q=1.0, *;q=0.5")
	return req, nil
}

func (s *CfsScraper) HandleResponse(j infra.Job, resp *http.Response, httpErr error) (*infra.Job, error) {
	scraperLog := log.WithField("job", j.Id).WithField("url", resp.Request.URL.String())
	// if we have an http error just return it
	if httpErr != nil {
		return nil, httpErr
	}

	var meta JobMetadata
	err := json.Unmarshal([]byte(j.MetaData), &meta)
	if err != nil {
		return nil, err
	}

	var reader io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(resp.Body)
		defer reader.Close()
	default:
		reader = resp.Body
	}

	var calls []*ServiceCall
	err = json.NewDecoder(reader).Decode(&calls)
	if err != nil {
		scraperLog.Error(err)
		return nil, err
	}

	for _, c := range calls {
		c.AfterParse()
	}

	start := time.Now()
	err = s.repo.StoreCalls(calls)
	if err != nil {
		scraperLog.Error(err)
		return nil, err
	}
	scraperLog.Debugf("Stored calls in %dms", time.Since(start).Milliseconds())

	// if we have at least LimitPerPage, we need to check for another page
	if len(calls) == LimitPerPage {
		newUrl := stripQueryParam(resp.Request.URL.String(), "$limit")
		newUrl = stripQueryParam(newUrl, "$offset")
		meta.Offset += LimitPerPage
		newMeta, err := json.Marshal(meta)
		scraperLog.Debug("New Job Metadata" + string(newMeta))
		if err != nil {
			return nil, err
		}
		return &infra.Job{
			Url:         newUrl,
			ScraperName: ScraperName,
			MetaData:    string(newMeta),
		}, nil
	} else {
		scraperLog.Debugf("num calls under 1000 %d", len(calls))
	}

	return nil, err
}

func stripQueryParam(inURL, stripKey string) string {
	u, err := url.Parse(inURL)
	if err != nil {
		return inURL
	}
	q := u.Query()
	q.Del(stripKey)
	u.RawQuery = q.Encode()
	return u.String()
}
