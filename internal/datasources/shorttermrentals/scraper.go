package shorttermrentals

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/codefornola/nolabase/internal/infra"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	ScraperName              = "SHORT-TERM-RENTALS"
	MergedVacationRentalsURL = "https://data.nola.gov/resource/en36-xvxg.json"
	LimitPerPage             = 1000
)

type JobMetadata struct {
	Limit  int
	Offset int
}

type STRScraper struct {
	repo *Repo
}

func NewScraper() *STRScraper {
	return &STRScraper{}
}

func (s *STRScraper) Configure(pool *pgxpool.Pool) (err error) {
	s.repo = NewRepo(pool)
	return err
}

func (s *STRScraper) EnqueueJobs() (jobs []infra.Job, err error) {
	metab, err := json.Marshal(&JobMetadata{
		Limit:  LimitPerPage,
		Offset: 0,
	})
	if err != nil {
		return jobs, err
	}
	job := infra.Job{
		Url:         MergedVacationRentalsURL,
		ScraperName: ScraperName,
		MetaData:    string(metab),
	}
	jobs = append(jobs, job)
	return jobs, nil
}

func (s *STRScraper) MakeRequest(j infra.Job) (*http.Request, error) {
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

func (s *STRScraper) HandleResponse(j infra.Job, resp *http.Response, httpErr error) (*infra.Job, error) {
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

	start := time.Now()
	log.Println("Decoding")
	var permits []*Permit
	err = json.NewDecoder(reader).Decode(&permits)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	log.Printf("Decoded in %dms\n", time.Since(start).Milliseconds())

	log.Println("AfterParse")
	for _, p := range permits {
		p.AfterParse()
	}

	log.Println("StorePermits")
	start = time.Now()
	err = s.repo.StorePermits(permits)
	if err != nil {
		return nil, err
	}
	log.Printf("Stored permits in %dms", time.Since(start).Milliseconds())

	log.Println("rentals stored")

	// if we have at least LimitPerPage, we need to check for another page
	if len(permits) == LimitPerPage {
		newUrl := stripQueryParam(resp.Request.URL.String(), "$limit")
		newUrl = stripQueryParam(newUrl, "$offset")
		meta.Offset += LimitPerPage
		newMeta, err := json.Marshal(meta)
		fmt.Println(string(newMeta))
		if err != nil {
			return nil, err
		}
		return &infra.Job{
			Url:         newUrl,
			ScraperName: ScraperName,
			MetaData:    string(newMeta),
		}, nil
	} else {
		fmt.Printf("num permits under 1000 %d", len(permits))
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
