package controllers

import (
	//"flower_shop/models"

	"github.com/astaxie/beego"
)

type TuijianController struct {
	beego.Controller
}

func (this *TuijianController) Get() {
	this.TplName = "tuijian.html"
}
