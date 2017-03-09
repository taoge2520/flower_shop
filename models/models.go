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
	_, err = stmt.Exec(cname, uname, t, totle)
	if err != nil {
		return
	}
	return
}
