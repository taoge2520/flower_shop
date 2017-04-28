package controllers

import (
	"github.com/astaxie/beego"
)

type WarningController struct {
	beego.Controller
}

func (this *WarningController) Get() {
	this.Data["ErrorInfo"] = "您已经评论过了"
	this.TplName = "warning.html"
	//this.TplName = "test.html"
}
