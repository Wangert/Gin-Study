package main

/**
用Handle通用方法获取http请求
 */

import (
	"github.com/gin-gonic/gin"
	"fmt"
)

func main()  {
	
	engine := gin.Default()

	//GET请求
	//http://localhost:8080/hello?name=wangert
	engine.Handle("GET", "/hello", func(context *gin.Context) {
		//获取请求全路径
		path := context.FullPath()
		fmt.Println(path)
		//获取请求参数
		name := context.DefaultQuery("name", "wjt")
		fmt.Println(name)

		//输出
		context.Writer.Write([]byte("hello, " + name))
	})

	//POST请求
	//http://localhost:8080/login
	engine.Handle("POST", "/login", func(context *gin.Context) {
		fmt.Println(context.FullPath())
		//获取表单提交数据
		username := context.PostForm("username")
		password := context.PostForm("password")

		fmt.Println(username)
		fmt.Println(password)

		context.Writer.Write([]byte(username + "\tLogin"))
	})

	engine.Run()
}
