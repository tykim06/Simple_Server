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

func GetNewestRecodesQuary(table string, group string, ilo_id int64) string {
	return "select * from " + table + " where Id in (select max(Id) from " + table + " where ILO_Id = " + strconv.FormatInt(ilo_id, 10) + " group by " + table + "." + group + ")"
}

func HpDBGetNewestRecodes(table string, group string, ilo_id int64, target interface{}) error {
	var original interface{}

	switch target.(type) {
	case *[]models.Fan:
		original = target.(*[]models.Fan)
	case *[]models.Temperature:
		original = target.(*[]models.Temperature)
	case *[]models.Power:
		original = target.(*[]models.Power)
	}
	_, err := Dbm.Select(original, GetNewestRecodesQuary(table, group, ilo_id))

	return err
}

func GetFansHealth(ilo_id int64) string {
	var fans []models.Fan
	if err := HpDBGetNewestRecodes("Fan", "FanName", ilo_id, &fans); err != nil {
		log.Println(err)
		return "Warning"
	} else {
		for _, f := range fans {
			if f.FanStatus.Health != "OK" {
				return "Warning"
			}
		}
	}

	return "OK"
}

func GetTemperaturesHealth(ilo_id int64) string {
	var temperatures []models.Temperature
	if err := HpDBGetNewestRecodes("Temperature", "Name", ilo_id, &temperatures); err != nil {
		log.Println(err)
		return "Warning"
	} else {
		for _, t := range temperatures {
			if t.TemperatureStatus.Health != "OK" && t.TemperatureStatus.State != "Absent" {
				return "Warning"
			}
		}
	}

	return "OK"
}

func GetPowersHealth(ilo_id int64) string {
	var powers []models.Power
	if err := HpDBGetNewestRecodes("Power", "BayNumber", ilo_id, &powers); err != nil {
		log.Println(err)
		return "Warning"
	} else {
		for _, p := range powers {
			if p.PowerStatus.Health != "OK" {
				return "Warning"
			}
		}
	}

	return "OK"
}

func GetTotalHealth(ilo_id int64, table string) string {
	switch table {
	case "Fan":
	case "Temperature":
	case "Power":
	}
	return "OK"
}

func HpDBGetOverviewInfo(ilo_id int64) map[string]string {
	totalHealthMap := make(map[string]string)
	totalHealthMap["Fans"] = GetTotalHealth(ilo_id, "Fan")
	totalHealthMap["Temperatures"] = GetTotalHealth(ilo_id, "Temperature")
	totalHealthMap["Powers"] = GetTotalHealth(ilo_id, "Power")

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
			GetNewestRecodesQuary("System", "SerialNumber", ilo.Id)); err != nil {
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
