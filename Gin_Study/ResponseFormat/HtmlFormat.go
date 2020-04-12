package main

/**
返回html页面
 */

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
)

func main()  {

	engine := gin.Default()

	//设置html目录
	engine.LoadHTMLGlob("ResponseFormat/html/*")
	//设置静态资源
	//前者为请求路径，后者为本地路径
	engine.Static("/img", "ResponseFormat/img/")

	engine.GET("/html", func(context *gin.Context) {

		fmt.Println(context.FullPath())
		context.HTML(http.StatusOK, "index.html", nil)
	})

	//给html页面传数据
	engine.GET("/template", func(context *gin.Context) {

		name := "Wangert is good!"
		context.HTML(http.StatusOK, "template.html", gin.H{
			"name":name,
		})

	})

	//显示图片
	engine.GET("/image", func(context *gin.Context) {

		fmt.Println(context.FullPath())
		context.HTML(http.StatusOK, "image.html", nil)
	})

	engine.Run()
}
