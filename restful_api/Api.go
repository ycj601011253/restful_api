package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:666666@tcp(127.0.0.1)/golong?charset=utf8") //链接数据库
	if err != nil { //如果数据库连接错误，直接结束，并关闭数据库
		return
	}
	// 创建一个存储的库，如果这个库不存在
	db.Exec("create database if not exists golong")
	// 创建一个接受用户名和密码的表
	db.Exec("create table if not exists user_info (id int unsigned auto_increment primary key,username varchar(16),password varchar(16))")
	defer db.Close()
	r := gin.Default()
	v := r.Group("/")
	{
		v.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
		v.POST("/register", register) //注册用户
		v.POST("/login", login)       //登录账号
	}
	r.Run() // listen and serve on 0.0.0.0:8080
}

func register(c *gin.Context) {
	username := c.PostForm("username")
	db, err := sql.Open("mysql", "root:666666@tcp(127.0.0.1)/golong?charset=utf8") //链接数据库
	if err != nil { //如果数据库连接错误，直接结束，并关闭数据库
		return
	}
	row := db.QueryRow("select username from user_info where username=?;", username) // 查询用户名是否存在
	var receive string
	err = row.Scan(&receive) // 如果用户名不存在会接收一个错误，执行插入数据操作
	if err != nil {
		password := c.PostForm("password")
		db.Exec("insert into user_info set username=?,password=?;", username, password) // 执行插入数据库语句
		c.JSON(200, gin.H{
			"result": "success", // 返回成功
		})
	} else {
		c.JSON(200, gin.H{
			"result": "failed", // 返回失败
			"reason": "The user name already exists", // 失败原因是用户名已存在
		})
	}
	defer db.Close() // 关闭数据库
}

func login(c *gin.Context) {
	db, err := sql.Open("mysql", "root:666666@tcp(127.0.0.1)/golong?charset=utf8") //链接数据库
	if err != nil { //如果数据库连接错误，直接返回，并关闭数据库
		return
	}
	username := c.PostForm("username")
	password := c.PostForm("password")
	var id int
	// 从数据库查询
	row := db.QueryRow("select id,username from user_info where username=? and password=?;",username,password)
	err = row.Scan(&id,&username)
	if err != nil{ // 如果出现错误，代表未查询到内容，返回重新登录
		c.JSON(200,gin.H{
			"message":"Login failed, please enter again",
		})
	}else { // 否则代表查询到内容，返回成功和用户名和id
		c.JSON(200,gin.H{
			"message":"Login success",
			"id":id,
			"username":username,
		})
	}
	defer db.Close()
}
