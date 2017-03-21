package controllers

import (
	"flower_shop/models"
	"fmt"
	"strconv"

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

func (this *Manage_comController) Add() {
	name := this.Input().Get("Name")
	price := this.Input().Get("Price")
	p, _ := strconv.Atoi(price)
	count := this.Input().Get("Count")
	c, _ := strconv.Atoi(count)
	fmt.Println(name, p, c)
	err := models.Add_commdity(name, p, c)
	if err != nil {
		fmt.Println(err)
		this.Redirect("/manage", 302)
		return
	}
	this.Redirect("/manage", 302)
	return
}
func (this *Manage_comController) Upd() {
	str_id := this.Input().Get("Id")
	id, _ := strconv.Atoi(str_id)
	name := this.Input().Get("Name")
	price := this.Input().Get("Price")
	p, _ := strconv.Atoi(price)
	count := this.Input().Get("Count")
	c, _ := strconv.Atoi(count)
	fmt.Println(id, name, p, c)
	err := models.Upd_commdity(id, name, p, c)
	if err != nil {
		fmt.Println(err)
		this.Redirect("/manage", 302)
		return
	}
	this.Redirect("/manage", 302)
	return
}
func (this *Manage_comController) Del() {
	str_id := this.Input().Get("Id")
	id, _ := strconv.Atoi(str_id)
	fmt.Println(id)
	err := models.Del_commdity(id)
	if err != nil {
		fmt.Println(err)
		this.Redirect("/manage", 302)
		return
	}
	this.Redirect("/manage", 302)
	return
}
