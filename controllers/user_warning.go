package controllers

import (
	"github.com/astaxie/beego"
)

type User_warningController struct {
	beego.Controller
}

func (this *User_warningController) Get() {
	this.Data["ErrorInfo"] = "用户名已经存在"
	this.TplName = "user_warning.html"
	//this.TplName = "test.html"
}
