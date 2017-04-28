package controllers

import (
	"flower_shop/models"
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
)

type ShopcarController struct {
	beego.Controller
}

func (c *ShopcarController) Get() {
	if Login_user == "" {
		c.Redirect("/", 302)
		return
	}
	datas, err := models.Get_gouwuche_uname(Login_user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(datas)
	c.Data["Datas"] = datas
	c.TplName = "gouwuche.html"
}
func (this *ShopcarController) Post() {
	if Login_user == "" {
		this.Redirect("/", 302)
		return
	}
	t1 := this.Input().Get("Name")
	t2 := this.Input().Get("Pirce")
	price, _ := strconv.Atoi(t2)
	t3 := this.Input().Get("Count")
	count, _ := strconv.Atoi(t3)
	if count == 0 {
		count = 1
	}
	err := models.Insert_shopcar(t1, price, count, Login_user) //两项操作
	if err != nil {
		fmt.Println(err)
	}
	this.Redirect("/gouwuche", 302)
	return
}
func (this *ShopcarController) Del() {
	str_id := this.Input().Get("Id_car")
	id, _ := strconv.Atoi(str_id)
	err := models.Delete_shopcar_by_id(id)
	if err != nil {
		fmt.Println(err)
	}
	this.Redirect("gouwuche", 302)
	return
}
