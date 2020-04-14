package controller

import "github.com/gin-gonic/gin"

type HelloController struct {

}

//路由方法
func (helloController *HelloController) Router(engine *gin.Engine)  {
	engine.GET("/hello", helloController.Hello)
}

//解析/hello路由哦请求
func (helloController *HelloController) Hello(context *gin.Context)  {
	context.JSON(200, map[string]interface{}{
		"Code":200,
		"Data":context.FullPath(),
	})
}