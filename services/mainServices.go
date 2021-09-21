package services

import (
	bc "bobee/controllers/base"
	"html/template"
)

// TemplateAdmin template for admin
func TemplateAdmin(c bc.AdminController) {
	c.Data["xsrf_token"] = c.XSRFToken()
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.LayoutSections["Header"] = "Layout/header.html"
	c.LayoutSections["Sidebar"] = "Layout/sidebar.html"
	c.LayoutSections["Footer"] = "Layout/footer.html"
}
