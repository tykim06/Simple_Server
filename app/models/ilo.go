package models

import (
	"time"

	"github.com/coopernurse/gorp"
)

type Ilo struct {
	Id        int64  `db:"Id"`
	Host      string `db:"Host"`
	User      string `db:"User"`
	Pass      string `db:"Pass"`
	CreatedAt string `db:"CreatedAt"`
}

func (c *Ilo) PreInsert(_ gorp.SqlExecutor) error {
	c.CreatedAt = time.Now().Format("2017-07-21 03:48:28")
	return nil
}
