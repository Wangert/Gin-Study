package main

/**
POST请求绑定表单数据
 */

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"log"
)

type RegisterUser struct {
	Name string `form:"name"`
	Password string `form:"password"`
	Email string `form:"email"`
}


func main()  {

	engine := gin.Default()

	engine.POST("/SignIn", func(context *gin.Context) {
		fmt.Println(context.FullPath())

		var registerUser RegisterUser
		//这里传地址
		err := context.ShouldBind(&registerUser)
		if err != nil {
			log.Fatal(err.Error())
			return
		}

		fmt.Println(registerUser.Name)
		fmt.Println(registerUser.Password)
		fmt.Println(registerUser.Email)

		context.Writer.WriteString("Name:" + registerUser.Name + "|Password:" + registerUser.Password + "|Email:" + registerUser.Email)

	})

	engine.Run()
}
