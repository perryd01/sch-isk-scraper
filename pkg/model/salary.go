package model

type Salary struct {
	Min   int64      `json:"min,omitempty"`
	Max   int64      `json:"max,omitempty"`
	Value int64      `json:"value,omitempty"`
	Unit  SalaryUnit `json:"unit,omitempty"`
}

type SalaryUnit string

const (
	Hour  SalaryUnit = "hour"
	Month SalaryUnit = "month"
)
