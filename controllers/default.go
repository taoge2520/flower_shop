package controllers

import (
	"flower_shop/models"
	"strconv"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	Nologin = ""
	commodities, err := models.Getcommdity()
	if err != nil {
		return
	}
	i := 1
	for _, v := range commodities {
		str := strconv.Itoa(i)
		Describe := "Desc" + str
		c.Data[Describe] = v.Name
		Picture := "Pic" + str
		c.Data[Picture] = "static/img/" + v.Picture
		i++
	}
	c.Data["IsLogin"] = Login
	c.TplName = "home.html"
}
