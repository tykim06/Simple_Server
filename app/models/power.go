package models

import (
	"time"

	"github.com/coopernurse/gorp"
)

type Power struct {
	Id          int64  `db:"Id" json:"Id"`
	ILO_Id      int64  `db:"Ilo_id" json:"Ilo_id"`
	Model       string `db:"Model" json:"Model"`
	Name        string `db:"Name" json:"Name"`
	PowerOem    `json:"Oem"`
	PowerStatus `json:"Status"`
	CreatedAt   string `db:"CreatedAt"`
}

type PowerOem struct {
	PowerHp `json:"Hp"`
}

type PowerHp struct {
	BayNumber int `db:"BayNumber" json:"BayNumber"`
}

type PowerStatus struct {
	Health string `db:"Health" json:"Health"`
	State  string `db:"State" json:"State"`
}

func (c *Power) PreInsert(_ gorp.SqlExecutor) error {
	c.CreatedAt = time.Now().Format("2017-07-21 03:48:28")
	return nil
}

func GetPowersHealth(powers []Power) string {
	for _, p := range powers {
		if p.PowerStatus.Health != "OK" {
			return "Warning"
		}
	}
	return "OK"
}

type PowerJson struct {
	Powers []Power `json:"PowerSupplies"`
}
