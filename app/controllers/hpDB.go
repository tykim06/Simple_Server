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
			for i, ilo := range ilos {
				HttpGetState(GET_STATE_SYSTEM, ilo, &systems[i])
				Dbm.Insert(&systems[i])

				var fanJson models.FanJson
				HttpGetState(GET_STATE_FAN, ilo, &fanJson)
				for _, fan := range fanJson.Fans {
					Dbm.Insert(&fan)
				}

				var powerJson models.PowerJson
				HttpGetState(GET_STATE_POWER, ilo, &powerJson)
				for _, power := range powerJson.Powers {
					Dbm.Insert(&power)
				}

				var temperatureJson models.TemperatureJson
				HttpGetState(GET_STATE_TEMPERATURE, ilo, &temperatureJson)
				for _, temperature := range temperatureJson.Temperatures {
					Dbm.Insert(&temperature)
				}
			}
			time.Sleep(10 * time.Second)
		}
	}()
}

func HpDBGetNewestRecode(table string, group string, target interface{}) {
	quary := "select * from " + table + " where id in (select max(id) from " + table + " group by " + table + "." + group + ")"

	switch table {
	case "Power":
		if original, ok := target.(*[]models.Power); ok {
			if _, err := Dbm.Select(original, quary); err != nil {
				panic(err)
			}
			log.Println(original)
		}
	case "Temperature":
		if original, ok := target.(*[]models.Temperature); ok {
			if _, err := Dbm.Select(original, quary); err != nil {
				panic(err)
			}
			log.Println(original)
		}

	case "Fan":
		if original, ok := target.(*[]models.Fan); ok {
			if _, err := Dbm.Select(original, quary); err != nil {
				panic(err)
			}
			log.Println(original)
		}

	}
}
