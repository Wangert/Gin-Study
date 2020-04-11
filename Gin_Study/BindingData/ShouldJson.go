package main

/**
用Json提交数据
 */

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"log"
	"strconv"
)

type Person struct {
	Name string `form:"name"`
	Age int `form:"age"`
	Sex string `form:"sex"`
}

func main()  {

	engine := gin.Default()

	engine.POST("/addUser", func(context *gin.Context) {
		fmt.Println(context.FullPath())

		var user Person
		if err := context.BindJSON(&user); err != nil {
			log.Fatal(err.Error())
			return
		}

		fmt.Println(user.Name)
		fmt.Println(user.Age)
		fmt.Println(user.Sex)

		context.Writer.WriteString("Name:" + user.Name + "|Age:" + strconv.Itoa(user.Age) + "|Sex:" + user.Sex)

	})

	engine.Run()
}
