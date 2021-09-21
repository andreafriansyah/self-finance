package routers

import (
	"bobee/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/home", &controllers.BooksController{}, "get:GetData")
	beego.Router("/sign-in", &controllers.LoginController{}, "get:ShowLogin")
	beego.Router("/sign-in", &controllers.LoginController{}, "post:DoSignIn")
	beego.Router("/validate-login", &controllers.LoginController{}, "post:ValidateLogin")
	beego.Router("/sign-out", &controllers.LoginController{}, "get:SignOut")

}
