package main

import (
	"github.com/gin-gonic/gin"
	"CloundRestaurant/tool"
	"CloundRestaurant/controller"
)

func main()  {
	//读取配置文件
	cfg, err := tool.ParseConfig("./config/app.json")
	if err != nil {
		panic(err)
	}
	//配置xorm操作引擎
	_, err = tool.OrmEngine(cfg)
	if err != nil {
		panic(err)
	}

	app := gin.Default()
	//设置路由
	registerRouter(app)

	app.Run(cfg.AppHost + ":" + cfg.AppPort)
}

//路由设置
func registerRouter(router *gin.Engine)  {
	new(controller.HelloController).Router(router)
	new(controller.MemberController).Router(router)
}
