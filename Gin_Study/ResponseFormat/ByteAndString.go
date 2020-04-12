package main

/**
返回byte数组
 */

import (
	"github.com/gin-gonic/gin"
	"fmt"
)

func main()  {

	engine := gin.Default()

	engine.GET("/byte", func(context *gin.Context) {

		fmt.Println(context.FullPath())
		context.Writer.Write([]byte("byte"))
	})

	engine.GET("/string", func(context *gin.Context) {

		fmt.Println(context.FullPath())
		context.Writer.WriteString("string")
	})

	engine.Run()
}
