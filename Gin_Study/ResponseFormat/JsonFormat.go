package main

/**
返回json格式数据
 */

import (
	"github.com/gin-gonic/gin"
	"fmt"
)

type Response struct {
	Code int
	Message string
	Data interface{}
}


func main()  {

	engine := gin.Default()

	engine.GET("/json", func(context *gin.Context) {

		fmt.Println(context.FullPath())
		//返回json格式数据
		context.JSON(200, map[string]interface{}{
			"code":1,
			"message":"ok",
			"data":"json test",
		})
	})

	engine.GET("/jsonStruct", func(context *gin.Context) {

		fmt.Println(context.FullPath())
		//返回json结构体数据
		response := Response{1, "ok", "json struct test"}

		context.JSON(200, &response)

	})

	engine.Run()
}
