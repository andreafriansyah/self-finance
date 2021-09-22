package controllers

import (
	base "bobee/controllers/base"
	"bobee/services"
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
