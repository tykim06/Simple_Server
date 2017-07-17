package models

import "time"

type System struct {
	Id        uint64 `db:"Id"`
	Model     string `db:"Model"`
	Health    string `db:"Health"`
	CreatedAt string `db:"CreatedAt"`
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
	system.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
}
