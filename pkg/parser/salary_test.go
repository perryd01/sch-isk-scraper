package parser

import (
	"testing"
)

func TestSalaryValue(t *testing.T) {
	s1 := "Br. 1630 Ft/óra*"
	s2 := "Br. 2000 Ft/óra*"

	salary1, err := pSalary(s1)
	if err != nil || salary1 == nil {
		t.Error("salary1 err:", err)
		return
	}
	if salary1.Value != 1630 {
		t.Error("invalid salary1 value", salary1.Value)
	}

	salary2, err := pSalary(s2)
	if err != nil || salary2 == nil {
		t.Error("salary2 err:", err)
		return
	}
	if salary2.Value != 2000 {
		t.Error("invalid salary2 value", salary2.Value)
	}

}

func TestSalaryRange(t *testing.T) {
	s3 := "Br. 1750-2200 Ft/óra*"

	salary3, err := pSalary(s3)
	if err != nil || salary3 == nil {
		t.Error("salary3 err:", err)
		return
	}
	if salary3.Min != 1750 && salary3.Max != 2200 {
		t.Error("invalid salary3 range: ", salary3.Min, salary3.Max)
	}

}

func TestSalaryValueEng(t *testing.T) {
	s4 := "Gross 2000 Ft/hour*"

	salary4, err := pSalary(s4)
	if err != nil || salary4 == nil {
		t.Error("salary4 err:", err)
		return
	}
	if salary4.Value != 2000 {
		t.Error("invalid salary4 value", salary4.Value)
	}
}

func TestSalaryRangeMonth(t *testing.T) {
	s5 := "Br. 480.000-600.000 Ft/hó*"

	salary5, err := pSalary(s5)
	if err != nil || salary5 == nil {
		t.Error("salary5 err:", err)
		return
	}
	if salary5.Min != 480_000 && salary5.Max != 600_000 && salary5.Unit != "month" {
		t.Error("invalid salary5 range or unit: ", salary5.Min, salary5.Max)
	}

}
