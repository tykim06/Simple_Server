package models

import (
	"fmt"
	"time"

	"github.com/coopernurse/gorp"
)

type System struct {
	Id        int64  `db:"id"`
	Model     string `db:"model"`
	Health    string `db:"health"`
	Power     string `db:"power"`
	CreatedAt string `db:"createdAt"`
}

func (c *System) PreInsert(_ gorp.SqlExecutor) error {
	c.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	return nil
}

func (c System) String() string {
	return fmt.Sprintf("Systems(%d, %s, %s, %s, %s)", c.Id, c.Model, c.Health, c.Power, c.CreatedAt)
}
