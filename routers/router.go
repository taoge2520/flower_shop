package routers

import (
	"flower_shop/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/login", &controllers.LoginController{}) //使用基础路由
	beego.Router("/", &controllers.MainController{})
	beego.Router("/home_search", &controllers.Home_searchController{})
	beego.Router("/gouwuche", &controllers.ShopcarController{})
	beego.Router("/gouwuche/del", &controllers.ShopcarController{}, "post:Del")

	beego.Router("/regist", &controllers.RegistController{})
	beego.Router("/warning", &controllers.WarningController{})
	beego.Router("/user_warning", &controllers.User_warningController{})
	beego.Router("/commodity/:id", &controllers.CommodityController{})
	beego.Router("/user", &controllers.UserController{})
	beego.Router("/order", &controllers.OrderController{})
	beego.Router("/tuijian", &controllers.TuijianController{})
	beego.Router("/root", &controllers.RootController{})

	beego.Router("/manage", &controllers.Manage_comController{})
	beego.Router("/manage_com/add", &controllers.Manage_comController{}, "post:Add")
	beego.Router("/manage_com/upd", &controllers.Manage_comController{}, "post:Upd")
	beego.Router("/manage_com/del", &controllers.Manage_comController{}, "post:Del")

	beego.Router("/manage_user", &controllers.Manage_userController{})
	beego.Router("/manage_user/add", &controllers.Manage_userController{}, "post:Add")
	//beego.Router("/manage_user/upd", &controllers.Manage_userController{}, "post:Upd")
	beego.Router("/manage_user/del", &controllers.Manage_userController{}, "post:Del")
}
