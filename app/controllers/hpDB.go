package controllers

import (
	"ilo/app/models"
	"log"
	"strconv"
	"time"
)

func InitHpDB() {
	go func() {
		log.Println("Start HpDB")
		var ilos []models.Ilo

		for {
			if _, err := HpDBGetIlo(&ilos); err != nil {
				log.Println(err)
			} else {
				for _, ilo := range ilos {
					go InsertCurrentState(ilo)
				}
			}

			time.Sleep(20 * time.Second)
		}
	}()
}

func InsertCurrentState(ilo models.Ilo) {
	var system models.System
	if err := HttpGetState(ilo, &system); err != nil {
		log.Println(err)
	} else {
		system.ILO_Id = ilo.Id
		system.ILO_Host = ilo.Host
		Dbm.Insert(&system)
	}

	time.Sleep(1 * time.Second)

	var fanTemperatureJson models.FanTemperatureJson
	if err := HttpGetState(ilo, &fanTemperatureJson); err != nil {
		log.Println(err)
	} else {
		for _, fan := range fanTemperatureJson.Fans {
			fan.ILO_Id = ilo.Id
			Dbm.Insert(&fan)
		}
		for _, temperature := range fanTemperatureJson.Temperatures {
			temperature.ILO_Id = ilo.Id
			Dbm.Insert(&temperature)
		}
	}

	time.Sleep(1 * time.Second)

	var powerJson models.PowerJson
	if err := HttpGetState(ilo, &powerJson); err != nil {
		log.Println(err)
	} else {
		for _, power := range powerJson.Powers {
			power.ILO_Id = ilo.Id
			Dbm.Insert(&power)
		}
	}
}

func HpDBGetIlo(ilos *[]models.Ilo) ([]interface{}, error) {
	*ilos = nil
	return Dbm.Select(ilos, `select * from Ilo`)
}

func HpDBGetNewestSystem(ilo_id int64, system *models.System) {
	quary := "select * from System where Id in (select max(Id) from System where ILO_Id = " + strconv.FormatInt(ilo_id, 10) + " group by System.SerialNumber)"
	if err := Dbm.SelectOne(system, quary); err != nil {
		log.Println(err)
	}
}

func HpDBGetNewestRecode(table string, group string, ilo_id int64, target interface{}) {
	quary := "select * from " + table + " where Id in (select max(Id) from " + table + " where ILO_Id = " + strconv.FormatInt(ilo_id, 10) + " group by " + table + "." + group + ")"

	var original interface{}
	var ok bool

	switch target.(type) {
	case *[]models.Power:
		original, ok = target.(*[]models.Power)
	case *[]models.Temperature:
		original, ok = target.(*[]models.Temperature)
	case *[]models.Fan:
		original, ok = target.(*[]models.Fan)
	}
	if !ok {
		log.Println("Type Assertion Failed")
	} else {
		if _, err := Dbm.Select(original, quary); err != nil {
			log.Println(err)
		}
	}
}
