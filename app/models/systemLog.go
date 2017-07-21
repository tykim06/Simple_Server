package models

import (
	"time"

	"github.com/coopernurse/gorp"
)

type SystemLog struct {
	Id              int64     `db:"Id"`
	ILO_Id          int64     `db:"ILO_Id"`
	Created         time.Time `db:"Created" json:"Created"`
	EntryType       string    `db:"EntryType" json:"EntryType"`
	Message         string    `db:"Message" json:"Message"`
	Name            string    `db:"Name" json:"Name"`
	Number          int       `db:"Number" json:"Number"`
	SystemLogOem    `db:"Oem" json:"Oem"`
	OemRecordFormat string    `db:"OemRecordFormat" json:"OemRecordFormat"`
	RecordID        int       `db:"RecordId" json:"RecordId"`
	Severity        string    `db:"Severity" json:"Severity"`
	CreatedAt       time.Time `db:"CreatedAt"`
	Page            int
}

type SystemLogOem struct {
	SystemLogHp `db:"Hp" json:"Hp"`
}

type SystemLogHp struct {
	Class   int       `db:"Class" json:"Class"`
	Code    int       `db:"Code" json:"Code"`
	Type    string    `db:"Type" json:"Type"`
	Updated time.Time `db:"Updated" json:"Updated"`
}

func (c *SystemLog) PreInsert(_ gorp.SqlExecutor) error {
	c.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	return nil
}

type SystemLogJson struct {
	SystemLogs []SystemLog `json:"Items"`
	Total      int         `json:"Total"`
}
