package controllers

import (
	"flower_shop/models"
	"fmt"

	"github.com/astaxie/beego"
)

type RegistController struct {
	beego.Controller
}

func (this *RegistController) Get() {
	this.TplName = "regist.html"
}

func (this *RegistController) Post() {
	uname := this.Input().Get("uname")
	pwd := this.Input().Get("pwd")
	phone := this.Input().Get("phone")
	fmt.Println(uname, pwd, phone)
	if uname == "" || pwd == "" {
		this.Redirect("/", 302)
		return
	}
	_, err := models.CheckSame(uname)
	if err == nil {
		this.Redirect("/warning", 302)
		return
	}
	err = models.Insert(uname, pwd, phone)
	if err != nil {
		fmt.Println(err)
		this.Redirect("/login", 302)
		return
	}
	this.Redirect("/", 302)

	return
}

// 验证用户名及密码
/*if uname == beego.AppConfig.String("adminName") &&
	pwd == beego.AppConfig.String("adminPass") {
	maxAge := 0
	if autoLogin {
		maxAge = 1<<31 - 1
	}

	this.Ctx.SetCookie("uname", uname, maxAge, "/")
	this.Ctx.SetCookie("pwd", pwd, maxAge, "/")
}

this.Redirect("/", 302)
return*/
