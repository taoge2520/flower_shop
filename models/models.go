package models

import (
	"database/sql"
	"fmt"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

var (
	DB *sql.DB
)

type Home_comm struct {
	Name    string
	Picture string
}
type order_datas struct {
	Id        int
	Comm_name string
	User      string
	Count     int
	Prices    int
}
type Comm_datas struct {
	Id         int
	Name       string
	Price      int
	Count      int
	Sell_count int
}

func init() {
	conn := beego.AppConfig.String("Conn")
	DB, _ = sql.Open("mysql", conn)
	DB.Ping()

}
func CheckSame(name string) (id int, err error) {
	err = DB.QueryRow("select id from user where uname=?", name).Scan(&id)
	if err != nil {
		return
	}
	return
}
func Get_phone(uname string) (phone string, err error) {
	err = DB.QueryRow("select phone from user where uname=?", uname).Scan(&phone)
	if err != nil {
		return
	}
	return
}
func Upd_phone(uname string, p string) (err error) {
	stmt, err := DB.Prepare("update user set phone=? where uname=?")
	if err != nil {
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(p, uname)
	if err != nil {
		return
	}
	return
}
func CheckInput(name string, pwd string) (id int, err error) {
	err = DB.QueryRow("select  id from user where uname=? and pwd=?", name, pwd).Scan(&id)
	if err != nil {
		return
	}
	return
}

func Insert(name string, pwd string, phone string) (err error) {
	stmt, err := DB.Prepare("insert user set uname=?,pwd=?,phone=?")
	if err != nil {
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(name, pwd, phone)
	if err != nil {
		return
	}
	return
}
func Getcommdity() (datas []Home_comm, err error) {
	rows, err := DB.Query(`select name ,picture from commodity limit 0,8`)
	if err != nil {
		return
	}
	var value Home_comm
	for rows.Next() {
		err = rows.Scan(&value.Name, &value.Picture)

		if err != nil {
			return
		}
		datas = append(datas, value)

	}
	return
}
func Get_comm_np(id int) (name string, price int, comms string) {
	err := DB.QueryRow("select name ,price ,comment_id from commodity where id =?", id).Scan(&name, &price, &comms)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}
func Get_comment(id int) (user string, str string, err error) {
	err = DB.QueryRow("select user,content from comment where id =?", id).Scan(&user, &str)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}
func Add_commodity(cname string, uname string, t int, totle int) (err error) {
	//insert into order_com (comm_name,user_name,count_c,totle_c) values("123","test1",10,20)
	stmt, err := DB.Prepare("insert into order_com (comm_name,user_name,count_c,totle_c) values(?,?,?,?)")
	if err != nil {
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(cname, uname, t, totle)
	if err != nil {
		return
	}
	return
}
func Get_order(user string) (s []order_datas, err error) {
	rows, err := DB.Query(`select id ,comm_name,user_name,count_c,totle_c from order_com where user_name=?`, user)
	if err != nil {
		return
	}

	for rows.Next() {
		var t order_datas
		err = rows.Scan(&t.Id, &t.Comm_name, &t.User, &t.Count, &t.Prices)

		if err != nil {
			return
		}
		s = append(s, t)

	}
	return
}

func Get_allcomm() (s []Comm_datas, err error) {
	rows, err := DB.Query(`select id ,name,price,count,sell_count from commodity`)
	if err != nil {
		return
	}

	for rows.Next() {
		var t Comm_datas
		err = rows.Scan(&t.Id, &t.Name, &t.Price, &t.Count, &t.Sell_count)

		if err != nil {
			return
		}
		s = append(s, t)

	}
	return
}

func Check_comment(id int, str string, user string) (err error) {
	var t int
	err = DB.QueryRow("select id from comment where id =? and user=?", id, user).Scan(&t)
	if err != nil {
		stmt, _ := DB.Prepare("insert comment set user=?,content=?")
		defer stmt.Close()
		_, err = stmt.Exec(user, str)
		if err != nil {
			fmt.Println(err)
		}
	}
	stmt, err := DB.Prepare("update comment set content=? where id=?")
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(str, t)
	if err != nil {
		fmt.Println(err)
	}
	return
}
