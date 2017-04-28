package controllers

import (
	"flower_shop/models"
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
)

type OrderController struct {
	beego.Controller
}

func (this *OrderController) Get() {
	orders, err := models.Get_order(Login_user)
	if err != nil {
		fmt.Println(err)
	}

	this.Data["Order_datas"] = orders
	this.TplName = "order.html"
}
func (this *OrderController) Post() {
	str_id := this.Input().Get("Orderid")
	content := this.Input().Get("Content")
	id, _ := strconv.Atoi(str_id)
	cname, err := models.Get_comm_name(id)
	if err != nil {
		this.Redirect("/order", 302)
		return
	}
	commodityid, err := models.Get_id_byname(cname) //获取商品id
	if err != nil {
		this.Redirect("/order", 302)
		return
	}
	err = models.Check_comment(Login_user, commodityid, id)
	if err != nil {
		err = models.Insert_comment(Login_user, content, commodityid, id)
		fmt.Println(err)
		this.Redirect("/order", 302)
		return
	}
	fmt.Println("aready commented!")
	this.Data["Aready"] = "您已经评论过了"
	this.Redirect("/warning", 302)
	return
}
