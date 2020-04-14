package controller

import (
	"github.com/gin-gonic/gin"
	"CloundRestaurant/service"
)

type MemberController struct {

}

func (memberController * MemberController) Router(engine *gin.Engine) {
	engine.GET("/api/sendSms", memberController.sendSmsCode)
}

//http://localhost:8888/api/sendSms?phone=123456789
func (memberController * MemberController) sendSmsCode(context *gin.Context)  {

	//获取手机号
	phone, exist := context.GetQuery("phone")
	if !exist {
		context.JSON(200, map[string]interface{}{
			"Code":0,
			"Message":"param is not exist.",
		})
	}

	//调用服务层
	memberService := service.MemberService{}
	memberService.SendSms(phone)

}