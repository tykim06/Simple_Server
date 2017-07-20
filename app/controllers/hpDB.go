package controllers

import (
	"ilo/app/models"
	"log"
	"time"
)

var (
	ilos    []models.Ilo
	systems []models.System
)

func InitHpDB() {
	if _, err := Dbm.Select(&ilos, `select * from Ilo`); err != nil {
		panic(err)
	}
	log.Println(len(ilos), "iLO Info Found")

	go func() {
		systems = make([]models.System, len(ilos))

		for {
			for i, _ := range ilos {
				go InsertCurrentState(i)
				time.Sleep(10 * time.Second)
			}
		}
	}()
}

func InsertCurrentState(i int) {
	if err := HttpGetState(GET_STATE_SYSTEM, ilos[i], &systems[i]); err == nil {
		systems[i].ILO_Id = ilos[i].Id
		Dbm.Insert(&systems[i])
	}

	var fanJson models.FanJson
	if err := HttpGetState(GET_STATE_FAN, ilos[i], &fanJson); err == nil {
		for _, fan := range fanJson.Fans {
			fan.ILO_Id = ilos[i].Id
			Dbm.Insert(&fan)
		}
	}

	var powerJson models.PowerJson
	if err := HttpGetState(GET_STATE_POWER, ilos[i], &powerJson); err == nil {
		for _, power := range powerJson.Powers {
			power.ILO_Id = ilos[i].Id
			Dbm.Insert(&power)
		}
	}

	var temperatureJson models.TemperatureJson
	if err := HttpGetState(GET_STATE_TEMPERATURE, ilos[i], &temperatureJson); err == nil {
		for _, temperature := range temperatureJson.Temperatures {
			temperature.ILO_Id = ilos[i].Id
			Dbm.Insert(&temperature)
		}
	}
}

func HpDBGetNewestRecode(table string, group string, target interface{}) {
	quary := "select * from " + table + " where id in (select max(id) from " + table + " group by " + table + "." + group + ")"

	var original interface{}
	var ok bool

	switch table {
	case "Power":
		original, ok = target.(*[]models.Power)
	case "Temperature":
		original, ok = target.(*[]models.Temperature)
	case "Fan":
		original, ok = target.(*[]models.Fan)
	}
	if !ok {
		log.Println("Type Assertion Failed")
		return
	}
	if _, err := Dbm.Select(original, quary); err != nil {
		log.Println("DB Select Error", err)
	}
}
