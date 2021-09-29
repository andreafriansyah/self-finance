package controllers

import (
	base "bobee/controllers/base"
	"bobee/services"
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
