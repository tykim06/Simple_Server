package controllers

import (
	"ilo/app/models"
	"log"
	"time"
)

var (
	ilos []*models.Ilo
)

func InitHpDB() {
	results, err := Dbm.Select(models.Ilo{}, `select * from Ilo`)
	if err != nil {
		panic(err)
	}

	for _, r := range results {
		ilos = append(ilos, r.(*models.Ilo))
	}
	log.Println(ilos)

	go func() {
		system := models.System{}
		systemJson := models.SystemJson{}
		for {
			for _, ilo := range ilos {
				getState(GET_STATE_SYSTEM, *ilo, &systemJson)
				systemJson.JsonToDB(&system)
				Dbm.Insert(&system)
			}
			time.Sleep(3 * time.Second)
		}
	}()
}
