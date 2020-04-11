package main

/**
GET请求绑定请求值
 */

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"log"
)

type User struct {
	Name string `form:"name"`
	Year string `form:"year"`
}

func main()  {

	engine := gin.Default()

	//http://localhost:8080/hello?name=wangert&year=2020
	engine.GET("/hello", func(context *gin.Context) {
		fmt.Println(context.FullPath())

		var user User
		//绑定对应值
		err := context.ShouldBindQuery(&user)
		if err != nil {
			log.Fatal(err.Error())
			return
		}

		fmt.Println(user.Name)
		fmt.Println(user.Year)

		context.Writer.Write([]byte("hello, " + user.Name))
	})

	engine.Run()
}
