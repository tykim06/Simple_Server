package controllers

import (
	"ilo/app/models"
	"ilo/app/routes"
	"log"

	"github.com/revel/revel"
)

type Monitor struct {
	GorpController
}

func (c Monitor) Index() revel.Result {
	var ilos []models.Ilo
	HpDBGetIlo(&ilos)

	systems := make([]models.System, len(ilos))
	for i, ilo := range ilos {
		HpDBGetNewestSystem(ilo.Id, &systems[i])
	}
	return c.Render(ilos, systems)
}

func (c Monitor) AddiLOForm() revel.Result {
	return c.Render()
}

func (c Monitor) AddiLO(ilo models.Ilo) revel.Result {
	if err := c.Txn.Insert(&ilo); err != nil {
		log.Println("Add iLO Failed")
	}
	return c.Redirect(routes.Monitor.Index())
}

func (c Monitor) Overview(ilo_id int64) revel.Result {
	totalHealthMap := map[string]string{
		"Fans":         "OK",
		"Temperatures": "OK",
		"Powers":       "OK",
	}

	var fans []models.Fan
	HpDBGetNewestRecode("Fan", "FanName", ilo_id, &fans)
	for _, f := range fans {
		if f.FanStatus.Health != "OK" {
			totalHealthMap["Fans"] = "Warning"
		}
	}

	var temperatures []models.Temperature
	HpDBGetNewestRecode("Temperature", "Name", ilo_id, &temperatures)
	for _, t := range temperatures {
		if t.TemperatureStatus.Health != "OK" && t.TemperatureStatus.State != "Absent" {
			totalHealthMap["Temperatures"] = "Warning"
		}
	}

	var powers []models.Power
	HpDBGetNewestRecode("Power", "BayNumber", ilo_id, &powers)
	for _, p := range powers {
		if p.PowerStatus.Health != "OK" {
			totalHealthMap["Powers"] = "Warning"
		}
	}

	return c.Render(ilo_id, totalHealthMap)
}

func (c Monitor) Fans(ilo_id int64) revel.Result {
	var fans []models.Fan
	HpDBGetNewestRecode("Fan", "FanName", ilo_id, &fans)
	return c.Render(ilo_id, fans)
}
func (c Monitor) Powers(ilo_id int64) revel.Result {
	var powers []models.Power
	HpDBGetNewestRecode("Power", "BayNumber", ilo_id, &powers)
	return c.Render(ilo_id, powers)
}
func (c Monitor) Temperatures(ilo_id int64) revel.Result {
	var temperatures []models.Temperature
	HpDBGetNewestRecode("Temperature", "Name", ilo_id, &temperatures)
	return c.Render(ilo_id, temperatures)
}

func (c Monitor) EventLog(ilo_id int64, pageNumber int) revel.Result {
	// eventJson := &models.EventLogJson{}
	// eventJson.Page = pageNumber
	// err := HttpGetState(ilos[id], eventJson)
	// if err != nil {
	// 	panic(err)
	// }
	// pageInfo := []int{}
	// for j, i := 1, eventJson.Total; i > 0; i -= 40 {
	// 	pageInfo = append(pageInfo, j)
	// 	j++
	// }
	// return c.Render(id, eventJson, pageInfo)
	return c.Render(ilo_id)
}

func (c Monitor) SystemLog(ilo_id int64, pageNumber int) revel.Result {
	// systemJson := &models.SystemLogJson{}
	// systemJson.Page = pageNumber
	// err := HttpGetState(ilos[id], systemJson)
	// if err != nil {
	// 	panic(err)
	// }
	// pageInfo := []int{}
	// for j, i := 1, systemJson.Total; i > 0; i -= 40 {
	// 	pageInfo = append(pageInfo, j)
	// 	j++
	// }
	// return c.Render(id, systemJson, pageInfo)
	return c.Render(ilo_id)
}
