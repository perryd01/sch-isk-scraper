package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/perryd01/sch-isk/pkg/cli"
	"github.com/perryd01/sch-isk/pkg/parser"
	"github.com/perryd01/sch-isk/pkg/requests"
	"log"
)

var options cli.Options

func init() {
	flag.StringVar((*string)(&options.City), "city", "budapest", "city of job")
	flag.StringVar((*string)(&options.JobType), "jobtype", "fejleszto---tesztelo", "type of job")
	flag.StringVar(&options.JobName, "jobname", "", "name of searched job")
	flag.UintVar(&options.MinimumSalary, "minSalary", 0, "minimum salary")
	flag.BoolVar(&options.Full, "allDetail", false, "get all of the details of a job")
	options.BaseUrl = "https://schonherz.hu"
}

func main() {
	flag.Parse()

	params := requests.RequestParams{
		IsListing: false,
		JobID:     0,
		JobName:   "",
		City:      options.City,
		JobType:   options.JobType,
	}

	listing, err := requests.FetchListing(params)
	if err != nil {
		log.Fatal(err)
	}

	jobs, err := parser.ReadListing(listing, options.BaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	if options.Full {
		for i := 0; i < len(jobs); i++ {
			doc, err := requests.FetchSingleJob(&jobs[i])
			if err != nil {
				log.Fatal(err)
			}
			parser.ReadOneJob(doc, &jobs[i])
		}
	}

	bytes, err := json.Marshal(jobs)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bytes))
}
