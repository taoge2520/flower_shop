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
	if Login_user == "" {
		this.Redirect("/", 302)
		return

	}
	phone, address, postcode, err := models.Get_userdatas(Login_user)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(address, postcode)
	this.Data["Phone"] = phone
	this.Data["Address"] = address
	this.Data["Postcode"] = postcode
	this.Data["Username"] = Login_user
	this.TplName = "user.html"
}

func (this *UserController) Post() {
	p := this.Input().Get("Upd_phone")
	a := this.Input().Get("Upd_address")
	c := this.Input().Get("Upd_postcode")
	phone, address, postcode, err := models.Get_userdatas(Login_user)
	if err != nil {
		fmt.Println(err)
		return
	}
	if p == "" {
		p = phone
	}
	if a == "" {
		a = address
	}
	if c == "" {
		c = postcode
	}

	err = models.Upd_p_a_c(Login_user, p, a, c)
	if err != nil {
		fmt.Println(err)
		return
	}
	this.Redirect("user", 302)
	return
}
