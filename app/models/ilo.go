package models

import (
	"fmt"
	"log"
	"time"

	"github.com/go-gorp/gorp"
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

func (c *Ilo) PreInsert(_ gorp.SqlExecutor) error {
	log.Println("ilo preinsert")
	return nil
}

func (i Ilo) String() string {
	return fmt.Sprintf("Ilos(%d, %s, %s, %s, %s)", i.Id, i.Host, i.User, i.Pass, i.CreatedAt)
}
