package controller

import (
	"github.com/gin-gonic/gin"
	"CloundRestaurant/service"
	"CloundRestaurant/param"
	"CloundRestaurant/tool"
	"fmt"
)

type MemberController struct {

}

func (memberController * MemberController) Router(engine *gin.Engine) {
	engine.GET("/api/sendSms", memberController.sendSmsCode)
	engine.POST("/api/register", memberController.register)
	engine.POST("/api/login", memberController.login)
}

//http://localhost:8888/api/sendSms?phone=123456789
func (memberController * MemberController) sendSmsCode(context *gin.Context)  {

	//获取手机号
	phone, exist := context.GetQuery("phone")
	if !exist {
		tool.Fail(context, "param is not exist.")
	}

	//调用服务层
	memberService := service.MemberService{}
	memberService.SendSms(phone)
}

//注册处理函数
func (memberController *MemberController) register(context *gin.Context) {
	//定义注册参数对象
	var registerParam param.RegisterParam
	//解析参数
	jsonParse := &tool.JsonParse{}
	err := jsonParse.Decode(context.Request.Body, &registerParam)
	if err != nil {
		tool.Fail(context, "param is not decode")
		return
	}

	//执行注册服务，并获得注册用户信息
	memberService := &service.MemberService{}
	member := memberService.Register(registerParam)
	if member != nil {
		tool.Success(context, member)
		return
	}

	tool.Fail(context, "register failed")
}


//登录处理函数
func (memberController *MemberController) login(context *gin.Context)  {
	//定义登录参数对象
	var loginParam param.LoginParam
	//解析参数
	jsonParse := &tool.JsonParse{}
	err := jsonParse.Decode(context.Request.Body, &loginParam)

	fmt.Println(loginParam.Username)

	if err != nil {
		tool.Fail(context, "param is not decode")
		return
	}

	//执行登录服务，并获得登录用户
	memberService := &service.MemberService{}
	member := memberService.Login(loginParam)
	if member.Id != 0 {
		tool.Success(context, member)
		return
	}

	tool.Fail(context, "login failed")
}