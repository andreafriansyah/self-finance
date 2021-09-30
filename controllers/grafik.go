package controllers

import (
	base "bobee/controllers/base"
	"bobee/services"
	servicesModel "bobee/services/grafikServices"
	"log"
)

type GrafikController struct {
	base.AdminController
}

func (c *GrafikController) GetData() {
	c.Layout = "Layout/layout.html"
	c.TplName = "Pages/grafik/index.html"
	c.LayoutSections["Script"] = "Pages/grafik/script.tpl"
	services.TemplateAdmin(c.AdminController)
}

func (c *GrafikController) JSON() {
	tahun := c.GetString("tahun")

	income, outcome, err := servicesModel.GetDataByFilter(tahun)
	if err != nil {
		log.Println("[Error] grafikController.Json : ", err)
		c.ServeError(err)
		return
	}

	c.Data["json"] = map[string]interface{}{
		"income":  income,
		"outcome": outcome,
	}
	c.ServeJSON()
}
