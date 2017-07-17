package models

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
