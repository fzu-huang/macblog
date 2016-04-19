package dbutil

import (
	"database/sql"
	"fmt"
	"github.com/fzu-huang/macblog/model"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
)

var db *sql.DB

var DBURL = `root:123@tcp(localhost:3306)/bbs?charset=utf8`

func init() {
	db, err := sql.Open("mysql", DBURL)
	fmt.Println(db)
	checkErr(err)
	db.SetMaxOpenConns(2000)
	db.SetMaxIdleConns(1800)
	err = db.Ping()
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func DBLogCheck(loguser model.LogUser) (bool, model.UserStatus) {
	username := template.HTMLEscapeString(loguser.Name)
	password := template.HTMLEscapeString(loguser.Pwd)

	db, err := sql.Open("mysql", DBURL)
	checkErr(err)
	defer db.Close()
	rows, err := db.Query("SELECT usermsg.name, usermsg.sex, usermsg.authority FROM usermsg where usermsg.name=? and usermsg.password=?", username, password)
	fmt.Println(err)
	checkErr(err)

	defer rows.Close()
	luser := model.UserStatus{}
	if rows.Next() {
		err = rows.Scan(&luser.Name, &luser.Sex, &luser.Authority)
		checkErr(err)
		luser.Online = true
		fmt.Println(luser)
		return true, luser
	}
	return false, luser
}

func DBRegCheck(reguser model.RegistUser) (bool, string) {
	username := template.HTMLEscapeString(reguser.Name)
	password := template.HTMLEscapeString(reguser.Password)
	sex := template.HTMLEscapeString(reguser.Sex)
	email := template.HTMLEscapeString(reguser.Email)
	authority := `会员`

	if DBUserExitCheck(username) {
		return false, "user  exist!!!"
	}

	db, err := sql.Open("mysql", DBURL)
	checkErr(err)
	defer db.Close()
	stmt, err := db.Prepare("insert into usermsg (name,password,sex,email,authority) values (?,?,?,?,?)")
	checkErr(err)
	result, err := stmt.Exec(username, password, sex, email, authority)
	checkErr(err)
	if err != nil {
		return false, err.Error()
	}
	fmt.Println(result)
	return true, ""
}

func DBUserExitCheck(username string) bool {
	db, err := sql.Open("mysql", DBURL)
	checkErr(err)
	defer db.Close()
	rows, err := db.Query("SELECT usermsg.name, usermsg.authority FROM usermsg where usermsg.name=?", username)
	fmt.Println(err)
	checkErr(err)

	var name, authority string
	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(&name, &authority)
		checkErr(err)
		fmt.Println(name, authority)
		return true
	}
	return false
}
