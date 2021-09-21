package controllers

import (
	"bobee/models"
	"bobee/models/shared"
	"bobee/services"

	"html/template"
	"log"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"github.com/jinzhu/gorm"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) ShowLogin() {
	sessionCurrentUser := c.GetSession("current_user")
	if sessionCurrentUser != nil {
		c.Redirect("/", 302)
	}

	c.Data["xsrf_token"] = c.XSRFToken()
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.TplName = "Auth/login.html"
}

// Validate Login with ajax
func (c *LoginController) ValidateLogin() {
	errorMessage := []map[string]interface{}{}
	errorStatus := false
	email := c.GetString("email")
	password := c.GetString("password")

	//validate login
	_, err := services.SignIn(email, password)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			errorMessage = append(errorMessage, map[string]interface{}{"message": "Email or password invalid"})
			errorStatus = true
		} else {
			c.Ctx.Output.SetStatus(400)
			response := shared.JSONResponse{
				Code:   400,
				Status: "Failed",
			}
			log.Println("[Error] AuthController.ValidateLogin BadRequest : ", err)
			c.Data["json"] = response
			c.ServeJSON()
			return
		}
	}

	if errorStatus {
		response := shared.JSONResponseValidate{
			Code:    400,
			Status:  "Failed",
			Message: errorMessage,
		}
		c.Data["json"] = response
		c.ServeJSON()
		return
	}

	response := shared.JSONResponse{
		Code:    200,
		Status:  "Success",
		Message: "Success sign in",
	}
	c.Data["json"] = response
	c.ServeJSON()
	return
}

func (c *LoginController) DoSignIn() {
	//template
	c.Data["xsrf_token"] = c.XSRFToken()
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.TplName = "Auth/login.html"

	form := models.User{}
	if err := c.ParseForm(&form); err != nil {
		log.Println("[Error] AuthController.DoSignIn : ", err)
		c.Abort("500")
	}
	c.Data["Form"] = form

	// check validate
	valid := validation.Validation{}
	valid.Required(form.Email, "Email").Message("is required")
	valid.Email(form.Email, "Email").Message("must email")
	valid.Required(form.Password, "Password").Message("is required")
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

	//check account
	dataUsers, err := services.SignIn(form.Email, form.Password)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			c.Data["HasErrors"] = true
			c.Data["Errors"] = map[string]string{"Login": "Email or password invalid"}
			return
		}
		log.Println("[Error] AuthController.SendEmailForgot : ", err)
		c.Abort("500")
	}

	c.SetSession("current_user", dataUsers)
	c.Redirect("/home", 302)
}

// SignOut destroy session
func (c *LoginController) SignOut() {
	c.DelSession("current_user")
	c.DestroySession()
	c.Redirect("/", 302)
}
