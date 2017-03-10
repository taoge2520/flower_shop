package controllers

import (
	"flower_shop/models"
	"fmt"

	"github.com/astaxie/beego"
)

type Manage_comController struct {
	beego.Controller
}

func (this *Manage_comController) Get() {
	all, err := models.Get_allcomm()
	if err != nil {
		fmt.Println(err)
	}
	this.Data["AllComm"] = all
	this.TplName = "manage_comm.html"
}
