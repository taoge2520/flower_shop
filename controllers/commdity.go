package controllers

import (
	"flower_shop/models"
	"fmt"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
)

type CommodityController struct {
	beego.Controller
}

func (this *CommodityController) Get() {
	str := this.Ctx.Input.Param(":id")
	h := strings.Split(str, "=")
	fmt.Println(h)
	id, _ := strconv.Atoi(h[1])

	name, price, picture := models.Get_comm_np(id) //商品名
	comments, err := models.Get_comment(id)
	if err != nil {
		return
	}
	if len(comments) == 0 {
		var t []models.Comment
		this.Data["Pic_show"] = "/static/img/" + picture
		this.Data["Comm_name"] = name
		this.Data["Comm_price"] = price
		this.Data["Comments"] = t
		this.Data["Page"] = h[1]
		this.TplName = "commdity.html"
		return
	}
	this.Data["Pic_show"] = "/static/img/" + picture
	this.Data["Comm_name"] = name
	this.Data["Comm_price"] = price
	this.Data["Comments"] = comments
	this.Data["Page"] = h[1]

	this.TplName = "commdity.html"
}

func (this *CommodityController) Post() {
	if Login_user == "" {
		this.Redirect("/", 302)
		return
	}
	str := this.Ctx.Input.Param(":id")
	h := strings.Split(str, "=")
	fmt.Println(h)
	id, _ := strconv.Atoi(h[1])

	name, price, totle_count := models.Get_comm_npc(id) //商品名
	count := this.Input().Get("Count")
	t, _ := strconv.Atoi(count)

	totle := t * price
	//默认的添加行为
	err := models.Add_commodity(name, Login_user, t, totle)
	if err != nil {
		fmt.Println(err)

	}
	new_count := totle_count - t
	err = models.Decrease(id, new_count)
	if err != nil {
		fmt.Println(err)
	}
	this.Redirect("/order", 302)
	return
}
