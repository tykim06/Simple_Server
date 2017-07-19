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

type FanJson struct {
	Fans []Fan `json:"Fans"`
}
