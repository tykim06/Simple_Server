package controllers

import (
	"ilo/app/models"
	"ilo/app/routes"
	"log"
	"time"

	"github.com/revel/revel"
)

type Monitor struct {
	GorpController
}

func (c Monitor) Index() revel.Result {
	return c.Render(ilos, systems)
}

func (c Monitor) AddiLOForm() revel.Result {
	return c.Render()
}

func (c Monitor) AddiLO(ilo models.Ilo) revel.Result {
	ilo.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	log.Println(ilo)
	err := c.Txn.Insert(&ilo)
	if err != nil {
		panic(err)
	}
	return c.Redirect(routes.Monitor.Index())
}

func (c Monitor) Overview(id int) revel.Result {
	subSystems := []string{"Fans", "Powers", "Temperatures"}
	states := []string{"OK", "OK", "OK"}
	fanJson := &models.FanJson{}
	err := HttpGetState(GET_STATE_FAN, *ilos[id], fanJson)
	if err != nil {
		panic(err)
	}
	for _, fan := range fanJson.Fans {
		if fan.Status.Health != "OK" {
			states[0] = "Warning"
			break
		}
	}

	powerJson := &models.PowerJson{}
	err = HttpGetState(GET_STATE_POWER, *ilos[id], powerJson)
	if err != nil {
		panic(err)
	}
	for _, power := range powerJson.PowerSupplies {
		if power.Oem.Hp.PowerSupplyStatus.State != "Ok" {
			states[1] = "Warning"
			break
		}
	}

	temperatureJson := &models.TemperatureJson{}
	err = HttpGetState(GET_STATE_TEMPERATURE, *ilos[id], temperatureJson)
	if err != nil {
		panic(err)
	}
	for _, temperature := range temperatureJson.Temperatures {
		if temperature.Status.Health != "OK" && temperature.Status.State != "Absent" {
			states[2] = "Warning"
			break
		}
	}

	return c.Render(id, subSystems, states)
}

func (c Monitor) Fans(id int) revel.Result {
	fanJson := &models.FanJson{}
	err := HttpGetState(GET_STATE_FAN, *ilos[id], fanJson)
	if err != nil {
		panic(err)
	}
	return c.Render(id, fanJson)
}
func (c Monitor) Powers(id int) revel.Result {
	powerJson := &models.PowerJson{}
	err := HttpGetState(GET_STATE_POWER, *ilos[id], powerJson)
	if err != nil {
		panic(err)
	}
	return c.Render(id, powerJson)
}
func (c Monitor) Temperatures(id int) revel.Result {
	temperatureJson := &models.TemperatureJson{}
	err := HttpGetState(GET_STATE_TEMPERATURE, *ilos[id], temperatureJson)
	if err != nil {
		panic(err)
	}
	return c.Render(id, temperatureJson)
}

func (c Monitor) EventLog(id int, pageNumber int) revel.Result {
	eventJson := &models.EventLogJson{}
	eventJson.Page = pageNumber
	err := HttpGetState(GET_STATE_EVENT_LOG, *ilos[id], eventJson)
	if err != nil {
		panic(err)
	}
	pageInfo := []int{}
	for j, i := 1, eventJson.Total; i > 0; i -= 40 {
		pageInfo = append(pageInfo, j)
		j++
	}
	return c.Render(id, eventJson, pageInfo)
}

func (c Monitor) SystemLog(id int, pageNumber int) revel.Result {
	systemJson := &models.SystemLogJson{}
	systemJson.Page = pageNumber
	err := HttpGetState(GET_STATE_SYSTEM_LOG, *ilos[id], systemJson)
	if err != nil {
		panic(err)
	}
	pageInfo := []int{}
	for j, i := 1, systemJson.Total; i > 0; i -= 40 {
		pageInfo = append(pageInfo, j)
		j++
	}
	return c.Render(id, systemJson, pageInfo)
}
