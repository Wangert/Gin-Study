package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
)

func main()  {

	engine := gin.Default()
	//全局使用中间件
	engine.Use(RequestInfos())

	engine.GET("/middleware", middlewareTestHandle)

	//单独为该请求设置中间件
	engine.GET("/print", PrintName(), func(context *gin.Context) {

	})

	engine.Run()

}

//打印请求信息的中间件
func RequestInfos() gin.HandlerFunc {
	return func(context *gin.Context) {
		fullPath := context.FullPath()
		method := context.Request.Method

		fmt.Println("Full Path:", fullPath)
		fmt.Println("Method:", method)
		fmt.Println("Before HttpStatus:", context.Writer.Status())

		//中断，先执行外部下一个业务请求，等请求结束，再执行以下内容
		context.Next()
		//状态码只有请求结束才能获得
		fmt.Println("After HttpStatus:", context.Writer.Status())
	}
}

//print请求专门的中间件
func PrintName() gin.HandlerFunc {
	return func(context *gin.Context) {
		fmt.Println("wangert")
	}
}

//中间件测试处理函数
func middlewareTestHandle(context *gin.Context) {

	fmt.Println("middlewareTestHandle start!")

	context.JSON(404, map[string]interface{}{
		"code":404,
		"Message":"ok",
		"Data":context.FullPath(),
	})
}
