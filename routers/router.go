package routers

import (
	"flower_shop/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/login", &controllers.LoginController{}) //使用基础路由
	beego.Router("/", &controllers.MainController{})
	beego.Router("/regist", &controllers.RegistController{})
	beego.Router("/warning", &controllers.WarningController{})
	beego.Router("/commodity", &controllers.CommodityController{})
	beego.Router("/user", &controllers.UserController{})
	beego.Router("/order", &controllers.OrderController{})
	beego.Router("/tuijian", &controllers.TuijianController{})
	beego.Router("/root", &controllers.RootController{})
	beego.Router("/manage", &controllers.Manage_comController{})
}
