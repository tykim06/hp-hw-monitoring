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
	ilos, systems := HpDBGetIndexInfo()
	return c.Render(ilos, systems)
}

func (c Monitor) AddiLOForm() revel.Result {
	return c.Render()
}

func (c Monitor) AddiLO(ilo models.Ilo) revel.Result {
	if err := c.Txn.Insert(&ilo); err != nil {
		log.Println(err)
	}
	return c.Redirect(routes.Monitor.Index())
}

func (c Monitor) Overview(ilo_id int64) revel.Result {
	totalHealthMap := HpDBGetOverviewInfo(ilo_id)

	return c.Render(ilo_id, totalHealthMap)
}

func (c Monitor) Fans(ilo_id int64) revel.Result {
	// var fans []models.Fan
	// if original, err := HpDBGetNewestRecodes("Fan", ilo_id); err != nil {
	// 	fans, _ := original.([]models.Fan)
	// }
	// return c.Render(ilo_id, fans)
	return c.Render(ilo_id)
}
func (c Monitor) Powers(ilo_id int64) revel.Result {
	// var powers []models.Power
	// if original, err := HpDBGetNewestRecodes("Power", ilo_id); err != nil {
	// 	powers, _ := original.([]models.Power)
	// }
	// return c.Render(ilo_id, powers)
	return c.Render(ilo_id)
}
func (c Monitor) Temperatures(ilo_id int64) revel.Result {
	// var temperatures []models.Temperature
	// HpDBGetNewestRecodes("Temperature", "Name", ilo_id, &temperatures)
	// return c.Render(ilo_id, temperatures)
	return c.Render(ilo_id)
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
