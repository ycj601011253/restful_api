package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:666666@tcp(127.0.0.1)/golong?charset=utf8") //链接数据库
	if err != nil { //如果数据库连接错误，直接结束，并关闭数据库
		return
	}
	db.Exec("create database if not exists golong")
	// 创建一个接受用户名和密码的表
	db.Exec("create table if not exists user_info (id int unsigned auto_increment primary key,username varchar(16),password varchar(16))")
	defer db.Close()
}
