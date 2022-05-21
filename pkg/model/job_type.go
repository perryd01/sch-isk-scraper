package model

import "encoding/json"

type JobTypeEnum string

const (
	Administrative JobTypeEnum = "adminisztrativ"
	Occasional     JobTypeEnum = "alkalmi---projekt-jellegu"
	Developer      JobTypeEnum = "fejleszto---tesztelo"
	Economy        JobTypeEnum = "gazdasagi"
	GrapicDesign   JobTypeEnum = "grafikus---designer"
	ITSupport      JobTypeEnum = "informatikai---support"
	Engineer       JobTypeEnum = "muszaki"
	Remote         JobTypeEnum = "otthonrol-vegezheto"
	UnderEightteen JobTypeEnum = "szaktudast-nem-igenylo-18-ev-alatt-"
	Foreign        JobTypeEnum = "jobs-for-foreign-students"
)

func (jt *JobTypeEnum) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}
	*jt = JobTypeEnum(s)
	return nil
}

func (jt *JobTypeEnum) MarshalJSON() ([]byte, error) {
	return []byte("\"" + string(*jt) + "\""), nil
}
