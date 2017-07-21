package models

import (
	"time"

	"github.com/coopernurse/gorp"
)

type Ilo struct {
	Id        int64     `db:"Id"`
	Host      string    `db:"Host"`
	User      string    `db:"User"`
	Pass      string    `db:"Pass"`
	CreatedAt time.Time `db:"CreatedAt"`
}

func (c *Ilo) PreInsert(_ gorp.SqlExecutor) error {
	c.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	return nil
}
