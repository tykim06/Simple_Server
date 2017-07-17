package models

import (
	"fmt"
	"time"
)

type Ilo struct {
	Id        int64  `db:"Id"`
	Host      string `db:"Host"`
	User      string `db:"User"`
	Pass      string `db:"Pass"`
	CreatedAt string `db:"CreatedAt"`
	//transient
	CreatedAtTime time.Time
}

func (i Ilo) String() string {
	return fmt.Sprintf("Ilos(%d, %s, %s, %s, %s)", i.Id, i.Host, i.User, i.Pass, i.CreatedAt)
}
