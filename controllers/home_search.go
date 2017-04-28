package controllers

import (
	"flower_shop/models"
	"fmt"

	"github.com/astaxie/beego"
)

type Home_searchController struct {
	beego.Controller
}

var temp string

func (c *Home_searchController) Get() {
	data, err := models.Getcommdity_by_name(temp)
	if err != nil {
		fmt.Println(err)
	}
	c.Data["Desc"] = data.Name
	c.Data["Pic"] = "static/img/" + data.Picture
	c.Data["IsLogin"] = Login
	c.Data["Page"] = data.Id
	c.TplName = "home_search.html"
}
func (this *Home_searchController) Post() {
	temp = this.Input().Get("Search")
	fmt.Println(temp)
	this.Redirect("/home_search", 302)
}
