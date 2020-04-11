package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"fmt"
)

func main() {

	engine := gin.Default()
	engine.GET("/hello", func(context *gin.Context) {
		fmt.Println("请求路径:", context.FullPath())
		context.Writer.Write([]byte("Hello World, Gin\n"))
	})

	if err := engine.Run(); err != nil {
		log.Fatal(err.Error())
	}

}
