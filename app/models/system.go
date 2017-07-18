package models

import (
	"fmt"
	"time"

	"github.com/coopernurse/gorp"
)

type System struct {
	Id        int64  `db:"id"`
	Model     string `db:"model"`
	Health    string `db:"health"`
	Power     string `db:"power"`
	CreatedAt string `db:"createdAt"`
}

func (c *System) PreInsert(_ gorp.SqlExecutor) error {
	c.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	return nil
}

func (c System) String() string {
	return fmt.Sprintf("Systems(%d, %s, %s, %s, %s)", c.Id, c.Model, c.Health, c.Power, c.CreatedAt)
}

type SystemJson struct {
	Bios struct {
		Current struct {
			VersionString string `json:"VersionString"`
		} `json:"Current"`
	} `json:"Bios"`
	Model      string `json:"Model"`
	Name       string `json:"Name"`
	Power      string `json:"Power"`
	Processors struct {
		Count           int    `json:"Count"`
		ProcessorFamily string `json:"ProcessorFamily"`
		Status          struct {
			HealthRollUp string `json:"HealthRollUp"`
		} `json:"Status"`
	} `json:"Processors"`
	SKU          string `json:"SKU"`
	SerialNumber string `json:"SerialNumber"`
	Status       struct {
		Health string `json:"Health"`
		State  string `json:"State"`
	} `json:"Status"`
}

func (c SystemJson) JsonToDB(system *System) {
	system.Model = c.Model
	system.Health = c.Processors.Status.HealthRollUp
	system.Power = c.Power
}
