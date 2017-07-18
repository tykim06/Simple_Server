package controllers

import (
	"ilo/app/models"
	"log"
	"time"
)

var (
	ilos    []*models.Ilo
	systems []*models.System
)

func InitHpDB() {
	results, err := Dbm.Select(models.Ilo{}, `select * from Ilo`)
	if err != nil {
		panic(err)
	}

	for _, r := range results {
		ilos = append(ilos, r.(*models.Ilo))
		systems = append(systems, &models.System{})
	}
	log.Println(len(ilos), "iLO Info Found")

	go func() {
		systemJson := models.SystemJson{}
		for {
			for i, ilo := range ilos {
				HttpGetState(GET_STATE_SYSTEM, *ilo, &systemJson)
				systemJson.JsonToDB(systems[i])
				Dbm.Insert(systems[i])
			}
			time.Sleep(10 * time.Second)
		}
	}()
}
