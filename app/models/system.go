package models

import (
	"time"

	"github.com/coopernurse/gorp"
)

type System struct {
	Id           int64  `db:"Id"`
	ILO_Id       int64  `db:"ILO_Id"`
	ILO_Host     string `db:"ILO_Host"`
	Bios         `db:"Bios" json:"Bios"`
	Model        string `db:"Model" json:"Model"`
	Name         string `db:"Name" json:"Name"`
	Power        string `db:"Power" json:"Power"`
	Processors   `db:"Processors" json:"Processors"`
	SKU          string `db:"SKU" json:"SKU"`
	SerialNumber string `db:"SerialNumber" json:"SerialNumber"`
	SystemStatus `db:"Status" json:"Status"`
	CreatedAt    string `db:"CreatedAt"`
}

type Bios struct {
	Current `db:"Current" json:"Current"`
}

type Current struct {
	VersionString string `db:"VersionString" json:"VersionString"`
}

type Processors struct {
	Count           int    `db:"Count" json:"Count"`
	ProcessorFamily string `db:"ProcessorFamily" json:"ProcessorFamily"`
	ProcessorStatus `db:"Status" json:"Status"`
}

type ProcessorStatus struct {
	HealthRollUp string `db:"HealthRollUp" json:"HealthRollUp"`
}

type SystemStatus struct {
	Health string `db:"Health" json:"Health"`
	State  string `db:"State" json:"State"`
}

func (c *System) PreInsert(_ gorp.SqlExecutor) error {
	c.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	return nil
}
