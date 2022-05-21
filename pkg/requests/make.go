package requests

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"

	"github.com/perryd01/sch-isk/pkg/model"
)

func FetchListing(params RequestParams) (string, error) {
	var exit = false
	var page string
	var counter = 1
	for exit != true {
		reqBody := fmt.Sprintf("{\"office\":\"%s\",\"type\":\"%s\",\"text\":\"\",\"current\":%d}", params.City, params.JobType, counter)
		req, err := http.NewRequest("POST", fmt.Sprintf("https://schonherz.hu/tobbdiakmunka/%s/%s", params.City, params.JobType), bytes.NewBuffer([]byte(reqBody)))
		if err != nil {
			return "", err
		}
		referer := fmt.Sprintf("https://schonherz.hu/diakmunkak/%s/%s", params.City, params.JobType)
		req.Header = map[string][]string{
			"accept":          {"*/*"},
			"accept-language": {"en-GB,en;q=0.9,en-US;q=0.8,hu;q=0.7"},
			"cache-control":   {"no-cache"},
			"content-type":    {"application/json; charset=UTF-8"},
			"sec-fetch-mode":  {"cors"},
			"sec-fetch-site":  {"same-origin"},
			"dnt":             {"1"},
			"origin":          {"https://schonherz.hu"},
			"pragma":          {"no-cache"},
			"referer":         {referer},
		}
		client := &http.Client{}
		res, err := client.Do(req)
		if err != nil {
			return "", err
		}
		defer func(Body io.ReadCloser) {
			_ = Body.Close()
		}(res.Body)
		if res.StatusCode != 200 {
			return "", err
		}

		b, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatalln(err)
		}

		page = page + string(b)
		counter++

		if len(string(b)) < 1 {
			exit = true
			return page, nil
		}

	}
	return page, nil
}

func FetchSingleJob(job *model.Job) (*goquery.Document, error) {
	resp, err := http.Get(job.Link)
	if err != nil {
		log.Fatalln(err)
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	return doc, nil
}

type RequestParams struct {
	IsListing bool              `json:"is_listing,omitempty"`
	JobID     uint64            `json:"job_id,omitempty"`
	JobName   string            `json:"job_name,omitempty"`
	City      model.CityEnum    `json:"city,omitempty"`
	JobType   model.JobTypeEnum `json:"job_type,omitempty"`
}
