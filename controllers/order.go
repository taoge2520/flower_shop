package controllers

import (
	"github.com/astaxie/beego"
)

type OrderController struct {
	beego.Controller
}

func (this *OrderController) Get() {
	this.TplName = "order.html"
}
