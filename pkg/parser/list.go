package parser

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/forPelevin/gomoji"
	"github.com/perryd01/sch-isk/pkg/model"
	"strings"
)

func ReadListing(html string, baseUrl string) ([]model.Job, error) {
	jobList := make([]model.Job, 0)

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return nil, err
	}
	doc.Find("body").Children().Each(func(i int, s *goquery.Selection) {

		link := s.Find("div.sdcard-outer > a")
		linkHref, _ := link.Attr("href")
		jobName := s.Find("div.text > h4 > a").Text()
		shortDescription := gomoji.RemoveEmojis(s.Find("div.text > span").Text())
		split := strings.Split(linkHref[1:], "/")
		codeName := strings.Split(split[3], "-")
		city := model.CityEnum(split[1])
		jobType := model.JobTypeEnum(split[2])

		tmpjob := model.Job{
			Name:         jobName,
			Description:  shortDescription,
			Code:         codeName[0],
			Expectations: "",
			Preference:   "",
			HoursPerWeek: "",
			WorkingPlace: "",
			Salary:       "",
			Link:         baseUrl + linkHref,
			City:         city,
			JobType:      jobType,
		}
		jobList = append(jobList, tmpjob)
	})
	return jobList, nil
}

func ReadOneJob(doc *goquery.Document, job *model.Job) {
	job.Description = gomoji.RemoveEmojis(doc.Find("#ad-details > div.row > div > span:nth-child(2)").Text())
	job.Expectations = doc.Find("#ad-details > div.row > div > span:nth-child(4)").Text()
	job.Preference = doc.Find("#ad-details > div.row > div > span:nth-child(6)").Text()
	job.HoursPerWeek = doc.Find("#ad-details > div.row > div > span:nth-child(8)").Text()
	job.WorkingPlace = doc.Find("#ad-details > div.row > div > span:nth-child(10)").Text()
	job.Salary = doc.Find("#ad-details > div.row > div > span:nth-child(12)").Text()
}
