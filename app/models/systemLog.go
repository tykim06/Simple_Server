package models

import (
	"time"
)

type SystemLogJson struct {
	Items []struct {
		Created   time.Time `json:"Created"`
		EntryType string    `json:"EntryType"`
		Message   string    `json:"Message"`
		Name      string    `json:"Name"`
		Number    int       `json:"Number"`
		Oem       struct {
			Hp struct {
				Class   int       `json:"Class"`
				Code    int       `json:"Code"`
				Type    string    `json:"Type"`
				Updated time.Time `json:"Updated"`
			} `json:"Hp"`
		} `json:"Oem"`
		OemRecordFormat string `json:"OemRecordFormat"`
		RecordID        int    `json:"RecordId"`
		Severity        string `json:"Severity"`
	} `json:"Items"`
	Total int `json:"Total"`
	Page  int
}
