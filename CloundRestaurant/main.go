package main

import (
	"github.com/gin-gonic/gin"
	"CloundRestaurant/tool"
	"CloundRestaurant/controller"
	"strings"
	"fmt"
	"net/http"
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

	//配置redis
	tool.InitRedis()

	app := gin.Default()

	//初始化session
	tool.InitSession(app)
	//设置全局跨域访问中间件
	app.Use(Cors())
	//设置路由
	registerRouter(app)

	app.Run(cfg.AppHost + ":" + cfg.AppPort)
}

//路由设置
func registerRouter(router *gin.Engine)  {
	new(controller.HelloController).Router(router)
	new(controller.MemberController).Router(router)
}

//跨域访问
func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		origin := context.Request.Header.Get("Origin")

		var headerKeys []string
		for key, _ := range context.Request.Header {
			headerKeys = append(headerKeys, key)
		}

		headerStr := strings.Join(headerKeys, ",")
		if headerStr != "" {
			headerStr = fmt.Sprint("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}

		if origin != "" {
			context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			// 设置允许访问所有域
			context.Header("Access-Control-Allow-Origin", "*")
			context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			context.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar")
			context.Header("Access-Control-Max-Age", "172800")
			context.Header("Access-Control-Allow-Credentials", "false")
			// 设置返回格式是json
			context.Set("content-type", "application/json")
		}

		if method == "OPTIONS" {
			context.JSON(http.StatusOK, "Option Request")
		}

		//继续处理请求
		context.Next()
	}
}
