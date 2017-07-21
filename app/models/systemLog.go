package models

import (
	"time"

	"github.com/coopernurse/gorp"
)

type SystemLog struct {
	Id              int64  `db:"Id"`
	ILO_Id          int64  `db:"ILO_Id"`
	Created         string `db:"Created" json:"Created"`
	EntryType       string `db:"EntryType" json:"EntryType"`
	Message         string `db:"Message" json:"Message"`
	Name            string `db:"Name" json:"Name"`
	Number          int    `db:"Number" json:"Number"`
	SystemLogOem    `db:"Oem" json:"Oem"`
	OemRecordFormat string `db:"OemRecordFormat" json:"OemRecordFormat"`
	RecordID        int    `db:"RecordId" json:"RecordId"`
	Severity        string `db:"Severity" json:"Severity"`
	CreatedAt       string `db:"CreatedAt"`
}

type SystemLogOem struct {
	SystemLogHp `db:"Hp" json:"Hp"`
}

type SystemLogHp struct {
	Class   int    `db:"Class" json:"Class"`
	Code    int    `db:"Code" json:"Code"`
	Type    string `db:"Type" json:"Type"`
	Updated string `db:"Updated" json:"Updated"`
}

func (c *SystemLog) PreInsert(_ gorp.SqlExecutor) error {
	c.CreatedAt = time.Now().Format("2017-07-21 03:48:28")
	return nil
}

type SystemLogJson struct {
	SystemLogs []SystemLog `json:"Items"`
	Total      int         `json:"Total"`
	Page       int
}
