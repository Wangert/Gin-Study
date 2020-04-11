package main

/**
将http请求分类
 */

import (
	"github.com/gin-gonic/gin"
	"fmt"
)


func main()  {

	engine := gin.Default()

	//http://localhost:8080/hello?name=wangert
	engine.GET("/hello", func(context *gin.Context) {
		fmt.Println(context.FullPath())

		name := context.Query("name")
		fmt.Println(name)

		context.Writer.Write([]byte("hello, " + name))
	})

	engine.POST("/login", func(context *gin.Context) {
		fmt.Println(context.FullPath())

		username, exist := context.GetPostForm("username")
		if exist {
			fmt.Println(username)
		}

		password, exist := context.GetPostForm("password")
		if exist {
			fmt.Println(password)
		}

		context.Writer.Write([]byte(username + " login"))
	})

	//:id表示变量值
	engine.DELETE("/user/:id", func(context *gin.Context) {
		id := context.Param("id")
		fmt.Println(id)

		context.Writer.Write([]byte("Delete " + id + " user"))
	})

	engine.Run()
}
