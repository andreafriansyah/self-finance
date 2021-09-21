package controllers

import (
	base "bobee/controllers/base"
	"bobee/services"
)

type BooksController struct {
	base.AdminController
}

func (this *BooksController) GetData() {
	this.Layout = "Layout/layout.html"
	this.TplName = "Pages/Book/index.html"
	services.TemplateAdmin(this.AdminController)
}

func (this *BooksController) PostData() {
	//
}
