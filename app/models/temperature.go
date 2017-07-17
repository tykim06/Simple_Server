package models

import (
	"fmt"
)

type Temperature struct {
	Id           int64  `db:"Id"`
	Temperature1 string `db:"Temperature1"`
	Temperature2 string `db:"Temperature2"`
	Temperature3 string `db:"Temperature3"`
}

type TemperatureJson struct {
	Temperatures []struct {
		Context                   string `json:"Context"`
		CurrentReading            int    `json:"CurrentReading"`
		LowerThresholdCritical    int    `json:"LowerThresholdCritical"`
		LowerThresholdNonCritical int    `json:"LowerThresholdNonCritical"`
		Name                      string `json:"Name"`
		Number                    int    `json:"Number"`
		Oem                       struct {
			Hp struct {
				LocationXmm int    `json:"LocationXmm"`
				LocationYmm int    `json:"LocationYmm"`
				Type        string `json:"Type"`
			} `json:"Hp"`
		} `json:"Oem"`
		Status struct {
			Health string `json:"Health"`
			State  string `json:"State"`
		} `json:"Status"`
		Units string `json:"Units"`
	} `json:"Temperatures"`
}

func (t Temperature) String() string {
	return fmt.Sprintf("Temperatures(%d, %s, %s, %s)", t.Id, t.Temperature1, t.Temperature2, t.Temperature3)
}
