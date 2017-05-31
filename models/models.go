package models

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

var (
	DB *sql.DB
)

type Home_comm struct {
	Id      int
	Name    string
	Picture string
}
type order_datas struct {
	Id        int
	Comm_name string
	User      string
	Count     int
	Prices    int
	Commented string
}
type Comm_datas struct {
	Id         int
	Name       string
	Price      int
	Count      int
	Sell_count int
}
type User_datas struct {
	Id    int
	Uname string
	Pwd   string
	Phone string
}
type Comment struct {
	User    string
	Content string
}

type Shopcars struct {
	Id           int
	Name         string
	Count        int
	Unit         int
	Uname        string
	Commodity_id int
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
func Get_userdatas(uname string) (phone string, addr string, postcode string, err error) {
	err = DB.QueryRow("select phone,addr,postcode from user where uname=?", uname).Scan(&phone, &addr, &postcode)
	if err != nil {
		return
	}
	return
}
func Upd_p_a_c(uname string, p string, a string, c string) (err error) {
	stmt, err := DB.Prepare("update user set phone=?,addr=?,postcode=? where uname=?")
	if err != nil {
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(p, a, c, uname)
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

func Insert(name string, pwd string, phone string, addr string, postcode string) (err error) {
	stmt, err := DB.Prepare("insert user set uname=?,pwd=?,phone=?,addr=?,postcode=?")
	if err != nil {
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(name, pwd, phone, addr, postcode)
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
func Getcommdity_by_name(name string) (datas Home_comm, err error) {
	err = DB.QueryRow(`select id ,name ,picture from commodity where name=?`, name).Scan(&datas.Id, &datas.Name, &datas.Picture)
	if err != nil {
		return
	}

	return
}
func Get_comm_np(id int) (name string, price int, pic string) {
	err := DB.QueryRow("select name ,price ,picture from commodity where id =?", id).Scan(&name, &price, &pic)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}
func Get_comm_npc(id int) (name string, price int, count int, sellcount int) {
	err := DB.QueryRow("select name ,price ,count ,sell_count from commodity where id =?", id).Scan(&name, &price, &count, &sellcount)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}
func Ranking() (names []string, sell_counts []int) {
	rows, err := DB.Query("select name ,sell_count from commodity order by sell_count desc limit 3")
	if err != nil {
		fmt.Println(err)
		return
	}
	for rows.Next() {
		var name string
		var t int
		err = rows.Scan(&name, &t)
		if err != nil {
			return
		}
		names = append(names, name)
		sell_counts = append(sell_counts, t)

	}
	return
}
func Decrease(id int, c int, sellcount int) (err error) {
	stmt, err := DB.Prepare("update commodity set count=?,sell_count=? where id=?")
	if err != nil {
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(c, sellcount, id)
	if err != nil {
		return
	}
	return
}
func Get_comment(id int) (comments []Comment, err error) {

	rows, err := DB.Query("select user,content from comment where commodity_id =?", id)
	if err != nil {
		return
	}
	for rows.Next() {
		var value Comment
		err = rows.Scan(&value.User, &value.Content)
		if err != nil {
			return
		}
		comments = append(comments, value)

	}
	return
}
func Add_order_com(cname string, uname string, t int, totle int) (err error) {
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
	rows, err := DB.Query(`select id ,comm_name,user_name,count_c,totle_c ,is_commented from order_com where user_name=?`, user)
	if err != nil {
		return
	}

	for rows.Next() {
		var t order_datas
		err = rows.Scan(&t.Id, &t.Comm_name, &t.User, &t.Count, &t.Prices, &t.Commented)

		if err != nil {
			return
		}
		s = append(s, t)

	}
	return
}

func Get_allcomm() (s []Comm_datas, err error) {
	rows, err := DB.Query(`select id ,name,price,count from commodity where is_delete=0`)
	if err != nil {
		return
	}

	for rows.Next() {
		var t Comm_datas
		err = rows.Scan(&t.Id, &t.Name, &t.Price, &t.Count)

		if err != nil {
			return
		}
		s = append(s, t)

	}
	return
}
func Get_alluser() (s []User_datas, err error) {
	rows, err := DB.Query(`select id ,uname,pwd,phone from user`)
	if err != nil {
		return
	}

	for rows.Next() {
		var t User_datas
		err = rows.Scan(&t.Id, &t.Uname, &t.Pwd, &t.Phone)

		if err != nil {
			return
		}
		s = append(s, t)

	}
	return
}
func Check_comment(user string, cid int, oid int) (err error) {
	var t int
	err = DB.QueryRow("select id from comment where user=? and commodity_id=? and orderid=? ", user, cid, oid).Scan(&t)
	if err != nil {
		return
	}
	return
}
func Get_comm_name(id int) (name string, err error) {
	err = DB.QueryRow("select comm_name from order_com where id =?", id).Scan(&name)
	if err != nil {
		return
	}
	return
}
func Get_id_byname(name string) (id int, err error) {
	err = DB.QueryRow("select id from commodity where name =?", name).Scan(&id)
	if err != nil {
		return
	}
	return
}

//Insert_comment(Login_user, content, commodityid, id)
func Insert_comment(user string, str string, cid int, id int) (err error) {
	stmt, err := DB.Prepare("insert into comment (user,content,commodity_id,orderid) values(?,?,?,?)")
	if err != nil {
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(user, str, cid, id)
	if err != nil {
		return
	}
	stmt1, err := DB.Prepare("update order_com set is_commented=? where id=?")
	if err != nil {
		return
	}
	defer stmt1.Close()
	_, err = stmt1.Exec("已评价", id)
	if err != nil {
		return
	}
	return
}
func Add_commdity(name string, price int, count int) (err error) {
	t := time.Now().Format("2006-01-02 15:04:05")
	stmt, err := DB.Prepare("insert into commodity (name,price,add_time,count) values(?,?,?,?)")
	if err != nil {
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(name, price, t, count)
	if err != nil {
		return
	}
	return
}
func Add_user(name string, pwd string, phone string) (err error) {
	//t := time.Now().Format("2006-01-02 15:04:05")
	stmt, err := DB.Prepare("insert into user (uname,pwd,phone) values(?,?,?)")
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
func Upd_commdity(id int, name string, price int, count int) (err error) {
	t := time.Now().Format("2006-01-02 15:04:05")
	stmt, err := DB.Prepare("update commodity set name=?,price=?,upd_time=?,count=? where id=?")
	if err != nil {
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(name, price, t, count, id)
	if err != nil {
		return
	}
	return
}
func Del_commdity(id int) (err error) {
	t := time.Now().Format("2006-01-02 15:04:05")
	stmt, err := DB.Prepare("update commodity set is_delete=?,upd_time=? where id=?")
	if err != nil {
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(1, t, id)
	if err != nil {
		return
	}
	return
}
func Del_user(id int) (err error) {
	stmt, err := DB.Prepare("delete from user where id=?")
	if err != nil {
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return
	}
	return
}
func Insert_shopcar(name string, price int, count int, uname string) (err error) {
	var t int
	err = DB.QueryRow("select id from commodity where name=?", name).Scan(&t)
	if err != nil {
		fmt.Println(err)
	}
	stmt, err := DB.Prepare("insert into gouwuche (name,unit,count,username,commodity_id) values(?,?,?,?,?)")
	if err != nil {
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(name, price, count, uname, t)
	if err != nil {
		return
	}
	return
}
func Get_gouwuche_uname(uname string) (v []Shopcars, err error) {
	rows, err := DB.Query(`select id,name ,count,unit,commodity_id from gouwuche where username=?`, uname)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var t Shopcars
		err = rows.Scan(&t.Id, &t.Name, &t.Count, &t.Unit, &t.Commodity_id)
		if err != nil {
			return
		}
		v = append(v, t)

	}
	return
}
func Delete_shopcar_by_id(id int) (err error) {
	_, err = DB.Exec("DELETE FROM gouwuche WHERE id=?", id)
	if err != nil {
		return
	}
	return
}
