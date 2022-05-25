package model

type Job struct {
	Name         string      `json:"name,omitempty"`
	Code         string      `json:"code"`
	Description  string      `json:"description,omitempty"`
	Expectations string      `json:"expectations,omitempty"`
	Preference   string      `json:"preference,omitempty"`
	HoursPerWeek string      `json:"hours_per_week,omitempty"`
	WorkingPlace string      `json:"working_place,omitempty"`
	Salary       Salary      `json:"salary,omitempty"`
	Link         string      `json:"link,omitempty"`
	City         CityEnum    `json:"city,omitempty"`
	JobType      JobTypeEnum `json:"job_type,omitempty"`
}
