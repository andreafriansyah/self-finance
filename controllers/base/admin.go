package controllers

import (
	"github.com/astaxie/beego"
)

type AdminController struct {
	beego.Controller
}

func (c *AdminController) Prepare() {
	sessionCurrentUser := c.GetSession("current_user")
	if sessionCurrentUser == nil {
		c.Redirect("/sign-in", 302)
	} else {
		c.Data["MainMenu"] = sessionCurrentUser
		c.LayoutSections = make(map[string]string)
	}
}

// Finish is called once the adminController method completes.
func (adminController *AdminController) Finish() {
	defer func() {
	}()
}

func (adminController *AdminController) ServeError(err error) {
	adminController.Data["json"] = struct {
		Error string `json:"Error"`
	}{err.Error()}
	adminController.Ctx.Output.SetStatus(500)
	adminController.ServeJSON()
}
