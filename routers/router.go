package routers

import (
	"bobee/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/sign-in", &controllers.LoginController{}, "get:ShowLogin")
	beego.Router("/sign-in", &controllers.LoginController{}, "post:DoSignIn")
	beego.Router("/validate-login", &controllers.LoginController{}, "post:ValidateLogin")
	beego.Router("/sign-out", &controllers.LoginController{}, "get:SignOut")

	beego.Router("/home", &controllers.HomeController{}, "get:GetData")

	beego.Router("/income", &controllers.PemasukanController{}, "get:GetIncome")
	beego.Router("/income/json", &controllers.PemasukanController{}, "get:JSON")

	beego.Router("/outcome", &controllers.PengeluaranController{}, "get:GetOutcome")
	beego.Router("/outcome/json", &controllers.PengeluaranController{}, "get:JSON")

}
