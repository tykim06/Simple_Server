package controllers

import (
	"ilo/app/models"
	"log"
	"strconv"
	"time"
)

var dbTableGroupMap = map[string]string{
	"Fan":         "FanName",
	"Power":       "BayNumber",
	"Temperature": "Name",
	"System":      "SerialNumber",
}

func InitHpDB() {
	go func() {
		log.Println("Start HpDB")
		var ilos []models.Ilo

		for {
			ilos = GetIlos()
			for _, ilo := range ilos {
				go InsertCurrentState(ilo)
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

func GetNewestRecodesQuary(table string, ilo_id int64) string {
	return "select * from " + table + " where Id in (select max(Id) from " + table + " where ILO_Id = " + strconv.FormatInt(ilo_id, 10) + " group by " + table + "." + dbTableGroupMap[table] + ")"
}

func GetFansTotalHealth(ilo_id int64) string {
	var fans []models.Fan
	if _, err := Dbm.Select(&fans, GetNewestRecodesQuary("Fan", ilo_id)); err != nil {
		log.Println(err)
		return "Warning"
	}

	return models.GetFansHealth(fans)
}

func GetTemperaturesTotalHealth(ilo_id int64) string {
	var temperatures []models.Temperature
	if _, err := Dbm.Select(&temperatures, GetNewestRecodesQuary("Temperature", ilo_id)); err != nil {
		log.Println(err)
		return "Warning"
	}

	return models.GetTemperaturesHealth(temperatures)
}

func GetPowersTotalHealth(ilo_id int64) string {
	var powers []models.Power
	if _, err := Dbm.Select(&powers, GetNewestRecodesQuary("Power", ilo_id)); err != nil {
		log.Println(err)
		return "Warning"
	}

	return models.GetPowersHealth(powers)
}

func HpDBGetOverviewInfo(ilo_id int64) map[string]string {
	totalHealthMap := make(map[string]string)
	totalHealthMap["Fans"] = GetFansTotalHealth(ilo_id)
	totalHealthMap["Temperatures"] = GetTemperaturesTotalHealth(ilo_id)
	totalHealthMap["Powers"] = GetPowersTotalHealth(ilo_id)

	return totalHealthMap
}

func GetIlos() []models.Ilo {
	var ilos []models.Ilo
	if _, err := Dbm.Select(&ilos, `select * from Ilo`); err != nil {
		log.Println(err)
	}
	return ilos
}

func GetNewestSystems(ilos []models.Ilo) []models.System {
	systems := make([]models.System, len(ilos))
	for i, ilo := range ilos {
		if err := Dbm.SelectOne(&systems[i],
			GetNewestRecodesQuary("System", ilo.Id)); err != nil {
			log.Println(err)
		}
	}
	return systems
}

func HpDBGetIndexInfo() ([]models.Ilo, []models.System) {
	ilos := GetIlos()
	systems := GetNewestSystems(ilos)
	return ilos, systems
}

func HpDBGetFansInfo(ilo_id int64) []models.Fan {
	var fans []models.Fan
	if _, err := Dbm.Select(&fans, GetNewestRecodesQuary("Fan", ilo_id)); err != nil {
		log.Println(err)
	}
	return fans
}

func HpDBGetPowersInfo(ilo_id int64) []models.Power {
	var powers []models.Power
	if _, err := Dbm.Select(&powers, GetNewestRecodesQuary("Power", ilo_id)); err != nil {
		log.Println(err)
	}
	return powers
}

func HpDBGetTemperaturesInfo(ilo_id int64) []models.Temperature {
	var temperature []models.Temperature
	if _, err := Dbm.Select(&temperature, GetNewestRecodesQuary("Temperature", ilo_id)); err != nil {
		log.Println(err)
	}
	return temperature
}
