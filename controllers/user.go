package controllers

import (
	"flower_shop/models"
	"fmt"

	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) Get() {
	phone, err := models.Get_phone(Login_user)
	if err != nil {
		fmt.Println(err)
		return
	}
	this.Data["Phone"] = phone
	this.Data["Username"] = Login_user
	this.TplName = "user.html"
}

func (this *UserController) Post() {
	p := this.Input().Get("Upd_phone")
	fmt.Println(p)
	err := models.Upd_phone(Login_user, p)
	if err != nil {
		fmt.Println(err)
		return
	}
	this.Redirect("user", 302)
	return
}
