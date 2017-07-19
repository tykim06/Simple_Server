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
	var fans []models.Fan
	var temperatures []models.Temperature
	return c.Render(id)
}

func (c Monitor) Fans(id int) revel.Result {
	var fans []models.Fan
	HpDBGetNewestRecode("Fan", "FanName", &fans)
	return c.Render(id, fans)
}
func (c Monitor) Powers(id int) revel.Result {
	var powers []models.Power
	HpDBGetNewestRecode("Power", "BayNumber", &powers)
	return c.Render(id, powers)
}
func (c Monitor) Temperatures(id int) revel.Result {
	var temperatures []models.Temperature
	HpDBGetNewestRecode("Temperature", "Name", &temperatures)
	return c.Render(id, temperatures)
}

func (c Monitor) EventLog(id int, pageNumber int) revel.Result {
	eventJson := &models.EventLogJson{}
	eventJson.Page = pageNumber
	err := HttpGetState(GET_STATE_EVENT_LOG, ilos[id], eventJson)
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
	err := HttpGetState(GET_STATE_SYSTEM_LOG, ilos[id], systemJson)
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
