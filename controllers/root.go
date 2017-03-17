package controllers

import (
	//"flower_shop/models"
	//"strconv"

	"github.com/astaxie/beego"
)

type RootController struct {
	beego.Controller
}

func (this *RootController) Get() {
	Nologin = ""
	this.TplName = "home1.html"
}

func (this *RootController) Post() {

}
