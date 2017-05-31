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
	ns, sc := models.Ranking()

	if len(ns) != 3 || len(sc) != 3 {
		c.TplName = "home.html"
		return
	}
	c.Data["First"] = ns[0]
	c.Data["Second"] = ns[1]
	c.Data["Third"] = ns[2]

	c.Data["First_count"] = sc[0]
	c.Data["Second_count"] = sc[1]
	c.Data["Third_count"] = sc[2]

	c.Data["IsLogin"] = Login
	c.TplName = "home.html"
}
