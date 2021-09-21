package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	sessionCurrentUser := c.GetSession("current_user")
	if sessionCurrentUser == nil {
		c.Redirect("/sign-in", 302)
	}
	c.Redirect("/home", 302)
}
