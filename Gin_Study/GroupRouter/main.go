package main

/**
路由分组(用于模块化开发)
 */

import (
	"github.com/gin-gonic/gin"
	"fmt"
)

func main()  {

	engine := gin.Default()
	//设置一个路由组
	routerGroup := engine.Group("/user")

	routerGroup.POST("/register", registerHandle)

	routerGroup.POST("/login", loginHandle)

	routerGroup.DELETE("/:id", deleteHandle)

	engine.Run()
}

//注册处理函数
func registerHandle(context *gin.Context) {
	fullPath := context.FullPath()
	fmt.Println(fullPath)

	context.Writer.WriteString(fullPath)
}
//登录处理函数
func loginHandle(context *gin.Context) {
	fullPath := context.FullPath()
	fmt.Println(fullPath)

	context.Writer.WriteString(fullPath)
}
//删除处理函数
func deleteHandle(context *gin.Context) {
	fullPath := context.FullPath()
	fmt.Println(fullPath)

	id := context.Param("id")
	context.Writer.WriteString(fullPath + id)
}

