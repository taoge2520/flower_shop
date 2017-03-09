package controllers

import (
	"github.com/astaxie/beego"
)

type WarningController struct {
	beego.Controller
}

func (this *WarningController) Get() {
	this.Data["ErrorInfo"] = "用户名已存在"
	this.TplName = "warning.html"
	//this.TplName = "test.html"
}
