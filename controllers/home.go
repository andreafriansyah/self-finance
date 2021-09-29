package controllers

import (
	base "bobee/controllers/base"
	"bobee/models"
	"bobee/services"
	servicesModel "bobee/services/saldoServices"
	"log"
)

type HomeController struct {
	base.AdminController
}

func (c *HomeController) GetData() {
	c.Layout = "Layout/layout.html"
	c.TplName = "Pages/Home/index.html"
	c.LayoutSections["Script"] = "Pages/Home/index.tpl"
	services.TemplateAdmin(c.AdminController)
}

func (c *HomeController) JSON() {
	resultData := []models.Saldo{}

	data, err := servicesModel.GetSaldo()
	if err != nil {
		log.Println("[Error] EmployeeController.Json : ", err)
		c.ServeError(err)
		return
	}
	resultData = data

	c.Data["json"] = map[string]interface{}{
		"data": resultData,
	}
	c.ServeJSON()
}
