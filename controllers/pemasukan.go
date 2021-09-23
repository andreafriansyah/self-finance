package controllers

import (
	base "bobee/controllers/base"
	"bobee/models"
	"bobee/services"

	servicesModel "bobee/services/incomeServices"
	"log"
	"strconv"

	"github.com/astaxie/beego/validation"
)

type PemasukanController struct {
	base.AdminController
}

func (c *PemasukanController) GetIncome() {
	c.Layout = "Layout/layout.html"
	c.TplName = "Pages/pemasukan/index.html"
	c.LayoutSections["Script"] = "Pages/pemasukan/script.tpl"
	services.TemplateAdmin(c.AdminController)
}

func (c *PemasukanController) JSON() {
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

func (c *PemasukanController) Tambah() {
	c.Layout = "Layout/layout.html"
	c.TplName = "Pages/pemasukan/tambah.html"
	c.LayoutSections["Script"] = "Pages/pemasukan/tambah.tpl"
	services.TemplateAdmin(c.AdminController)
}

func (c *PemasukanController) TambahData() {
	c.Layout = "Layout/layout.html"
	c.TplName = "Pages/pemasukan/tambah.html"
	c.LayoutSections["Script"] = "Pages/pemasukan/tambah.tpl"
	services.TemplateAdmin(c.AdminController)

	form := models.Finances{}
	if err := c.ParseForm(&form); err != nil {
		log.Println("[Error] InfoController.DoAdd BadRequest : ", err)
		c.Abort("500")
	}
	c.Data["Form"] = form

	//check validate
	//check validate
	valid := validation.Validation{}
	valid.Required(form.Jumlah, "Jumlah").Message("is required")
	valid.Required(form.AsalTujuan, "Asal").Message("is required")
	valid.Required(form.Keterangan, "Keterangan").Message("is required")

	errorMap := []string{}
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
			errorMap = append(errorMap, err.Key+" "+err.Message)
		}
		c.Data["HasErrors"] = true
		c.Data["Errors"] = errorMap
		return
	}

	//validate name
	_, err := servicesModel.Create(form)
	if err != nil {
		log.Println("[Error] InfoController.DoAdd BadRequest : ", err)
		c.Abort("500")
		return
	}

	c.Redirect("/income", 302)
}
