package model

import "encoding/json"

type CityEnum string

const (
	Budapest  CityEnum = "budapest"
	Debrecen  CityEnum = "debrecen"
	Kecskemet CityEnum = "kecskemet"
	Miskolc   CityEnum = "miskolc"
	Pecs      CityEnum = "pecs"
	Szeged    CityEnum = "szeged"
)

func (ce *CityEnum) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}
	*ce = CityEnum(s)
	return nil
}

func (ce *CityEnum) MarshalJSON() ([]byte, error) {
	return []byte("\"" + string(*ce) + "\""), nil
}
