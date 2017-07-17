package models

import (
	"fmt"
)

type Fan struct {
	Id   int64  `db:"Id"`
	Fan1 string `db:"Fan1"`
	Fan2 string `db:"Fan2"`
	Fan3 string `db:"Fan3"`
	Fan4 string `db:"Fan4"`
}

type FanJson struct {
	Fans []struct {
		CurrentReading int    `json:"CurrentReading"`
		FanName        string `json:"FanName"`
		Oem            struct {
			Hp struct {
				Location string `json:"Location"`
				Type     string `json:"Type"`
			} `json:"Hp"`
		} `json:"Oem"`
		Status struct {
			Health string `json:"Health"`
			State  string `json:"State"`
		} `json:"Status"`
		Units string `json:"Units"`
	} `json:"Fans"`
}

func (f Fan) String() string {
	return fmt.Sprintf("Fans(%d, %s, %s, %s, %s)", f.Id, f.Fan1, f.Fan2, f.Fan3, f.Fan4)
}
