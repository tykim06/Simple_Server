package models

import (
	"time"

	"github.com/coopernurse/gorp"
)

type EventLog struct {
	Id              int64     `db:"Id"`
	ILO_Id          int64     `db:"ILO_Id"`
	Created         time.Time `db:"Created" json:"Created"`
	EntryType       string    `db:"EntryType" json:"EntryType"`
	Message         string    `db:"Message" json:"Message"`
	Name            string    `db:"Name" json:"Name"`
	Number          int       `db:"Number" json:"Number"`
	EventLogOem     `db:"Oem" json:"Oem"`
	OemRecordFormat string    `db:"OemRecordFormat" json:"OemRecordFormat"`
	RecordID        int       `db:"RecordId" json:"RecordId"`
	Severity        string    `db:"Severity" json:"Severity"`
	Type            string    `db:"Type" json:"Type"`
	CreatedAt       time.Time `db:"CreatedAt"`
	Page            int
}

type EventLogOem struct {
	EventLogHp `db:"Hp" json:"Hp"`
}

type EventLogHp struct {
	Class   int       `db:"Class" json:"Class"`
	Code    int       `db:"Code" json:"Code"`
	Type    string    `db:"Type" json:"Type"`
	Updated time.Time `db:"Updated" json:"Updated"`
}

func (c *EventLog) PreInsert(_ gorp.SqlExecutor) error {
	c.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	return nil
}

type EventLogJson struct {
	EventLogs []EventLog `json:"Items"`
	Total     int        `json:"Total"`
}
