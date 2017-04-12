package controllers

import (
	"flower_shop/models"
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
)

type Manage_userController struct {
	beego.Controller
}

func (this *Manage_userController) Get() {
	all, err := models.Get_alluser()
	if err != nil {
		fmt.Println(err)
	}
	this.Data["Alluser"] = all
	this.TplName = "manage_user.html"
}

func (this *Manage_userController) Add() {
	name := this.Input().Get("Uname")
	pwd := this.Input().Get("Pwd")
	phone := this.Input().Get("Phone")
	fmt.Println(name, pwd, phone)
	err := models.Add_user(name, pwd, phone)
	if err != nil {
		fmt.Println(err)
		this.Redirect("/manage_user", 302)
		return
	}
	this.Redirect("/manage_user", 302)
	return
}

func (this *Manage_userController) Del() {
	str_id := this.Input().Get("Id")
	id, _ := strconv.Atoi(str_id)
	fmt.Println(id)
	err := models.Del_user(id)
	if err != nil {
		fmt.Println(err)
		this.Redirect("/manage_user", 302)
		return
	}
	this.Redirect("/manage_user", 302)
	return
}
