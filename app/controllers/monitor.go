package controllers

import (
	"ilo/app/models"
	"ilo/app/routes"
	"log"
	"time"

	"github.com/revel/revel"
)

var ilos []*models.Ilo

type Monitor struct {
	GorpController
	HttpController
}

func (c Monitor) Index() revel.Result {
	results, err := c.Txn.Select(models.Ilo{}, `select * from Ilo`)
	if err != nil {
		panic(err)
	}

	var temp_ilos []*models.Ilo
	systemJsons := make([]models.SystemJson, len(results))
	for i, r := range results {
		temp_ilos = append(temp_ilos, r.(*models.Ilo))
		c.getState(GET_STATE_SYSTEM, *temp_ilos[i], &systemJsons[i])
	}
	ilos = temp_ilos
	log.Println(ilos, systemJsons)
	return c.Render(ilos, systemJsons)
}

func (c Monitor) AddiLOForm() revel.Result {
	return c.Render()
}

func (c Monitor) AddiLO(ilo models.Ilo) revel.Result {
	ilo.CreatedAtTime = time.Now()
	ilo.CreatedAt = ilo.CreatedAtTime.Format("2006-01-02 15:04:05")
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
	err := c.getState(GET_STATE_FAN, *ilos[id], fanJson)
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
	err = c.getState(GET_STATE_POWER, *ilos[id], powerJson)
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
	err = c.getState(GET_STATE_TEMPERATURE, *ilos[id], temperatureJson)
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
	err := c.getState(GET_STATE_FAN, *ilos[id], fanJson)
	if err != nil {
		panic(err)
	}
	return c.Render(id, fanJson)
}
func (c Monitor) Powers(id int) revel.Result {
	powerJson := &models.PowerJson{}
	err := c.getState(GET_STATE_POWER, *ilos[id], powerJson)
	if err != nil {
		panic(err)
	}
	return c.Render(id, powerJson)
}
func (c Monitor) Temperatures(id int) revel.Result {
	temperatureJson := &models.TemperatureJson{}
	err := c.getState(GET_STATE_TEMPERATURE, *ilos[id], temperatureJson)
	if err != nil {
		panic(err)
	}
	return c.Render(id, temperatureJson)
}

func (c Monitor) EventLog(id int, pageNumber int) revel.Result {
	eventJson := &models.EventLogJson{}
	eventJson.Page = pageNumber
	err := c.getState(GET_STATE_EVENT_LOG, *ilos[id], eventJson)
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
	err := c.getState(GET_STATE_SYSTEM_LOG, *ilos[id], systemJson)
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
