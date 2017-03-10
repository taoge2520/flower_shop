package controllers

import (
	"flower_shop/models"

	"github.com/astaxie/beego"
)

var (
	Login      bool
	Login_root bool
)

type LoginController struct {
	beego.Controller
}

var (
	Login_user string
)

func (this *LoginController) Get() {
	Login = false
	Login_root = Login
	this.Data["IsLogin"] = Login
	this.TplName = "login.html"
}

func (this *LoginController) Post() {
	root_user := beego.AppConfig.String("admin")
	root_pwd := beego.AppConfig.String("pwd")
	uname := this.Input().Get("uname")
	pwd := this.Input().Get("pwd")
	if uname == root_user && pwd == root_pwd {

		Login_root = true
		this.Redirect("/root", 302)
		return

	}
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
