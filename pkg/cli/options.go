package cli

import "github.com/perryd01/sch-isk/pkg/model"

type Options struct {
	City          model.CityEnum
	JobType       model.JobTypeEnum
	JobName       string
	MinimumSalary uint
	Full          bool
	BaseUrl       string
}
