package controller

import (
	"github.com/gin-gonic/gin"
	"CloundRestaurant/service"
	"CloundRestaurant/param"
	"CloundRestaurant/tool"
	"fmt"
	"encoding/json"
	"CloundRestaurant/model"
	"strconv"
	"time"
)

type MemberController struct {

}

func (memberController * MemberController) Router(engine *gin.Engine) {
	engine.GET("/api/sendSms", memberController.sendSmsCode)
	engine.POST("/api/register", memberController.register)
	engine.POST("/api/login", memberController.login)
	engine.GET("/api/captcha", memberController.captcha)
	//头像上传
	engine.POST("/api/upload/avator", memberController.uploadAvator)
}

//头像上传
func (memberController * MemberController) uploadAvator(context *gin.Context)  {

	//解析上传数据
	userId := context.PostForm("user_id")
	fmt.Println(userId)
	file, err := context.FormFile("avator")
	if err != nil || userId == "" {
		tool.Fail(context, "avator failed")
		return
	}

	//判断是否登录
	memberSession := tool.SessionGet(context, "user_" + userId)
	if memberSession == nil {
		tool.Fail(context, "not login")
		return
	}

	//将session序列化数据，反序列化为Member对象
	var member model.Member
	json.Unmarshal(memberSession.([]byte), &member)

	fileName := "./uploadfile/" + strconv.FormatInt(time.Now().Unix(), 10) + file.Filename
	err = context.SaveUploadedFile(file, fileName)
	if err != nil {
		tool.Fail(context, "avator upload failed")
	}

	//http://localhost:8888/static/.../wangert.png
	//将保存在本地的文件路径，记录到数据库里
	memberService := service.MemberService{}
	path := memberService.UploadAvator(member.Id, fileName[1:])
	if path != "" {
		tool.Success(context, "http://localhost:8888" + path)
		return
	}

	tool.Fail(context, "avator insert db failed")

}

//生成验证码
func (memberController * MemberController) captcha(context *gin.Context)  {
	
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
		//序列化member信息
		memberSession, _ := json.Marshal(member)
		//将用户信息保存到session
		err = tool.SessionSet(context, "user_" + string(member.Id), memberSession)
		if err != nil {
			tool.Fail(context, "session set failed")
			return
		}

		tool.Success(context, member)
		return
	}

	tool.Fail(context, "login failed")
}