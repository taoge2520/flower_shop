package controllers

import (
	"flower_shop/models"
	"fmt"
	"net/smtp"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
)

type Mailto struct {
	User     string
	Password string
	Host     string
	To       string
	Mailtype string //类型
	Subject  string //标题
	Body     string //内容
}
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
	str_shopcar_id := this.Input().Get("Id_car")
	if str_shopcar_id != "" {
		shopcar_id, _ := strconv.Atoi(str_shopcar_id)
		err := models.Delete_shopcar_by_id(shopcar_id) //删除该条购物车记录
		if err != nil {
			fmt.Println(err)

		}
		str := this.Ctx.Input.Param(":id") //添加一条订单记录，更新库存
		h := strings.Split(str, "=")
		fmt.Println(h)
		id, _ := strconv.Atoi(h[1])

		name, price, totle_count, sell := models.Get_comm_npc(id) //商品名
		count := this.Input().Get("Count")

		t, _ := strconv.Atoi(count)
		if t == 0 {
			t = 1
		}
		remain := sell + t //卖出数量更新
		totle := t * price

		err = models.Add_order_com(name, Login_user, t, totle)
		if err != nil {
			fmt.Println(err)

		}
		new_count := totle_count - t

		err = models.Decrease(id, new_count, remain)
		if err != nil {
			fmt.Println(err)
		}
		this.Redirect("/gouwuche", 302)
		return
	}

	str := this.Ctx.Input.Param(":id")
	h := strings.Split(str, "=")
	fmt.Println(h)
	id, _ := strconv.Atoi(h[1])

	name, price, totle_count, sell := models.Get_comm_npc(id) //商品名
	if totle_count < 0 {
		Warnning("商品售罄，若有货可补，请更新库存数量!")
	}

	count := this.Input().Get("Count")

	t, _ := strconv.Atoi(count)
	if t == 0 {
		t = 1
	}
	remain := sell + t //卖出数量更新
	totle := t * price

	err := models.Add_order_com(name, Login_user, t, totle)
	if err != nil {
		fmt.Println(err)

	}
	new_count := totle_count - t
	if new_count < 50 {
		Warnning("库存数量低于设置阈值，请及时补货!")
	}
	fmt.Println(new_count, remain)
	err = models.Decrease(id, new_count, remain)
	if err != nil {
		fmt.Println(err)
	}
	this.Redirect("/order", 302)
	return
}
func Warnning(str string) {
	var t Mailto
	t.User = "18850046590@163.com"
	t.Password = "ziqun1993"
	t.Host = "smtp.163.com:25"
	t.To = "1853988263@qq.com"

	t.Subject = "告警邮件"
	t.Mailtype = "html"
	t.Body = `
    <html>
    <body>
    <h3>
    ` + str + `
    </h3>
    </body>
    </html>
    `
	err := t.SendMail()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("send mail success!")
	}
}
func (this Mailto) SendMail() (err error) {
	hp := strings.Split(this.Host, ":") //hp is :[smtp.163.com 25]
	//fmt.Print("hp is :", hp)
	auth := smtp.PlainAuth("", this.User, this.Password, hp[0]) //&{ 18030120049@163.com blt520.. smtp.163.com}
	//fmt.Print("auth is :", auth)
	var content_type string
	if this.Mailtype == "html" {
		content_type = "Content-Type: text/" + this.Mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + this.To + "\r\nFrom: " + this.User + "<" + this.User + ">\r\nSubject: " + this.Subject + "\r\n" + content_type + "\r\n\r\n" + this.Body)
	//fmt.Printf("%s", msg)
	//fmt.Println("to is :", this.To)
	send_to := strings.Split(this.To, ";")
	//fmt.Println("send_to is :", send_to)

	err = smtp.SendMail(this.Host, auth, this.User, send_to, msg)

	return err
}
