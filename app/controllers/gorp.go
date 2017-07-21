package controllers

import (
	//"code.google.com/p/go.crypto/bcrypt"
	"database/sql"
	"ilo/app/models"
	"log"

	"github.com/coopernurse/gorp"
	_ "github.com/go-sql-driver/mysql"
	r "github.com/revel/revel"
)

var (
	Dbm *gorp.DbMap
)

type GorpController struct {
	*r.Controller
	Txn *gorp.Transaction
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func InitDB() {
	db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/iLO")
	checkErr(err, "sql.Open failed")
	Dbm = &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}

	Dbm.AddTableWithName(models.Fan{}, "Fan").SetKeys(true, "Id")
	Dbm.AddTableWithName(models.Power{}, "Power").SetKeys(true, "Id")
	Dbm.AddTableWithName(models.Temperature{}, "Temperature").SetKeys(true, "Id")
	Dbm.AddTableWithName(models.System{}, "System").SetKeys(true, "Id")
	Dbm.AddTableWithName(models.Ilo{}, "Ilo").SetKeys(true, "Id")
	Dbm.AddTableWithName(models.EventLog{}, "EventLog").SetKeys(true, "Id")
	Dbm.AddTableWithName(models.EventLog{}, "SystemLog").SetKeys(true, "Id")
	err = Dbm.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")
	Dbm.TraceOn("[gorp]", r.INFO)

	log.Println("Success gorp initialize")
}

func (c *GorpController) Begin() r.Result {
	txn, err := Dbm.Begin()
	if err != nil {
		log.Println("pannic occured in Begin")
		panic(err)
	}
	c.Txn = txn
	return nil
}

func (c *GorpController) Commit() r.Result {
	if c.Txn == nil {
		return nil
	}
	if err := c.Txn.Commit(); err != nil && err != sql.ErrTxDone {
		log.Println("pannic occured in Commit")
		panic(err)
	}
	c.Txn = nil
	return nil
}

func (c *GorpController) Rollback() r.Result {
	if c.Txn == nil {
		return nil
	}
	if err := c.Txn.Rollback(); err != nil && err != sql.ErrTxDone {
		log.Println("pannic occured in Rollback")
		panic(err)
	}
	c.Txn = nil
	return nil
}
