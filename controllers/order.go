package controllers

import (
	"flower_shop/models"
	"fmt"
	//"strconv"

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
	fmt.Println(orders)
	this.Data["Order_datas"] = orders
	this.TplName = "order.html"
}
func (this *OrderController) Post() {
	str_id := this.Input().Get("Orderid")
	content := this.Input().Get("Content")
	fmt.Println(str_id, content)
	this.Redirect("/order", 302)
	//id, _ := strconv.Atoi(str_id)
	//_ = models.Check_comment(id, content, Login_user)
	//根据订单号获取商品名，拿着商品名去commodity更新相关评论id列表
	//this.Redirect("/order", 302)

}
