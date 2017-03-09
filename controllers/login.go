package controllers

import (
	"flower_shop/models"

	"github.com/astaxie/beego"
)

var (
	Login bool
)

type LoginController struct {
	beego.Controller
}

var (
	Login_user string
)

func (this *LoginController) Get() {
	Login = false
	this.TplName = "login.html"
}

func (this *LoginController) Post() {
	uname := this.Input().Get("uname")
	pwd := this.Input().Get("pwd")
	_, err := models.CheckInput(uname, pwd)
	if err != nil {
		Login = false
		this.Redirect("/", 302)
	}
	Login = true
	Login_user = uname
	this.Redirect("/", 302)
	return
}
