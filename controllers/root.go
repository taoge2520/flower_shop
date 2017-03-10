package controllers

import (
	//"flower_shop/models"
	//"strconv"

	"github.com/astaxie/beego"
)

type RootController struct {
	beego.Controller
}

func (c *RootController) Get() {

	c.TplName = "home1.html"
}
