package controllers

import (
	base "bobee/controllers/base"
	"bobee/models"
	"bobee/services"

	servicesModel "bobee/services/outcomeServices"
	"log"
	"strconv"
)

type PengeluaranController struct {
	base.AdminController
}

func (c *PengeluaranController) GetOutcome() {
	c.Layout = "Layout/layout.html"
	c.TplName = "Pages/pengeluaran/index.html"
	c.LayoutSections["Script"] = "Pages/pengeluaran/script.tpl"
	services.TemplateAdmin(c.AdminController)
}

func (c *PengeluaranController) JSON() {
	asaltujuan := c.GetString("asaltujuan")
	search := c.GetString("search[value]")
	orderColumn := c.GetString("order[0][column]")
	orderBy := c.GetString("columns[" + orderColumn + "][name]")
	orderDirection := c.GetString("order[0][dir]")
	start, _ := strconv.Atoi(c.GetString("start"))
	length, _ := strconv.Atoi(c.GetString("length"))

	resultData := []models.Finances{}
	var totalRows int

	data, totalData, err := servicesModel.GetDataByFilter(asaltujuan, search, orderBy, orderDirection, start, length)
	if err != nil {
		log.Println("[Error] EmployeeController.Json : ", err)
		c.ServeError(err)
		return
	}
	resultData = data
	totalRows = totalData

	c.Data["json"] = map[string]interface{}{
		"recordsTotal":    totalRows,
		"recordsFiltered": totalRows,
		"data":            resultData,
	}
	c.ServeJSON()
}
