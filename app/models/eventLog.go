package models

import (
	"fmt"
	"time"
)

type EventLog struct {
	Id          int64  `db:"Id"`
	Severity    string `db:"Severity"`
	Description string `db:"Description"`
	Date        string `db:"Date"`
}

type EventLogJson struct {
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
		Type            string `json:"Type"`
	} `json:"Items"`
	Total int `json:"Total"`
	Page  int
}

func (e EventLog) String() string {
	return fmt.Sprintf("EventLog(%d, %s, %s, %s)", e.Id, e.Severity, e.Description, e.Date)
}
