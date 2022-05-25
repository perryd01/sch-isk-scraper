package parser

import (
	"fmt"
	"github.com/perryd01/sch-isk/pkg/model"
	"regexp"
	"strconv"
)

func pSalary(s string) (*model.Salary, error) {
	salary := model.Salary{}

	match, err := regexp.MatchString("((Br.)\\s\\d*)", s)
	if err != nil {
		return nil, err
	}

	if match {
		var value string
		_, err := fmt.Sscanf(s, "Br. %s Ft/óra*", &value)
		if err != nil {
			return nil, err
		}
		salary.Value, err = strconv.ParseInt(value, 10, 64)
		if err != nil {
			return nil, err
		}
	}

	return &salary, nil
}
