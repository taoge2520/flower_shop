package routers

import (
	"hello/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/", &controllers.MainController{})
	beego.Router("/regist", &controllers.RegistController{})
	beego.Router("/warning", &controllers.WarningController{})
	beego.Router("/commodity", &controllers.CommodityController{})
	beego.Router("/user", &controllers.UserController{})
	beego.Router("/order", &controllers.OrderController{})

}
