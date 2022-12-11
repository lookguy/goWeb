package model

type User struct{
	Uid int
	Name string
	Phone string
}

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init(){
	db,_ =sql.Open("mysql","root:123456@tcp(127.0.0.1:3306)/chapter3")
}

func GetUser(uid int )(u User){
	err := DB.QueryRow("select uid, name, phone from `user` where uid = ?",uid).Scan(&u.Uid, &u.Name, &u.Phone)

	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return
	}
	
	return u
}