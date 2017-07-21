package models

import (
	"time"

	"github.com/coopernurse/gorp"
)

type Fan struct {
	Id             int64  `db:"Id"`
	ILO_Id         int64  `db:"ILO_Id"`
	CurrentReading int    `db:"CurrentReading" json:"CurrentReading"`
	FanName        string `db:"FanName" json:"FanName"`
	FanOem         `db:"Oem" json:"Oem"`
	FanStatus      `db:"Status" json:"Status"`
	Units          string `db:"Units" json:"Units"`
	CreatedAt      string `db:"CreatedAt"`
}

type FanOem struct {
	FanHp `db:"Hp" json:"Hp"`
}

type FanHp struct {
	Location string `db:"Location" json:"Location"`
	Type     string `db:"Type" json:"Type"`
}

type FanStatus struct {
	Health string `db:"Health" json:"Health"`
	State  string `db:"State" json:"State"`
}

func (c *Fan) PreInsert(_ gorp.SqlExecutor) error {
	c.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	return nil
}

func GetFansHealth(fans []Fan) string {
	for _, f := range fans {
		if f.FanStatus.Health != "OK" {
			return "Warning"
		}
	}
	return "OK"
}

type Temperature struct {
	Id                int64  `db:"Id"`
	ILO_Id            int64  `db:"ILO_Id"`
	Context           string `db:"Context" json:"Context"`
	CurrentReading    int    `db:"CurrentReading" json:"CurrentReading"`
	Name              string `db:"Name" json:"Name"`
	Number            int    `db:"Number" json:"Number"`
	TemperatureOem    `db:"Oem" json:"Oem"`
	TemperatureStatus `db:"Status" json:"Status"`
	Units             string `db:"Units" json:"Units"`
	CreatedAt         string `db:"CreatedAt"`
}

type TemperatureOem struct {
	TemperatureHp `db:"Hp" json:"Hp"`
}

type TemperatureHp struct {
	LocationXmm int    `db:"LocationXmm" json:"LocationXmm"`
	LocationYmm int    `db:"LocationYmm" json:"LocationYmm"`
	Type        string `db:"Type" json:"Type"`
}

type TemperatureStatus struct {
	Health string `db:"Health" json:"Health"`
	State  string `db:"State" json:"State"`
}

func (c *Temperature) PreInsert(_ gorp.SqlExecutor) error {
	c.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	return nil
}

func GetTemperaturesHealth(temperatures []Temperature) string {
	for _, t := range temperatures {
		if t.TemperatureStatus.Health != "OK" && t.TemperatureStatus.State != "Absent" {
			return "Warning"
		}
	}
	return "OK"
}

type FanTemperatureJson struct {
	Fans         []Fan         `json:"Fans"`
	Temperatures []Temperature `json:"Temperatures"`
}
