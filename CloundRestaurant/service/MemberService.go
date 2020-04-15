package service

import (
	"CloundRestaurant/model"
	"CloundRestaurant/param"
	"CloundRestaurant/dao"
	"time"
	"CloundRestaurant/tool"
)

type MemberService struct {

}

func (memberService *MemberService) SendSms(phone string) bool {
	return true
}

//用户注册服务
func (memberService *MemberService) Register(registerParam param.RegisterParam) *model.Member {

	//定义用户数据操作层对象
	memberDao := dao.MemberDao{tool.DBEngine}
	//判断是否用户名是否已经存在
	if !memberDao.Exsit(registerParam.Username) {
		return nil
	}

	//验证二次输入密码是否相同
	if registerParam.Password != registerParam.RePassword {
		return nil
	}

	member := model.Member{}
	member.Username = registerParam.Username
	member.Password = registerParam.Password
	member.RegisterTime = time.Now().Unix()

	if result := memberDao.AddMember(member); result == 0 {
		return nil
	}

	return &member
}

//用户登录服务
func (memberService *MemberService) Login(loginParam param.LoginParam) *model.Member {

	//验证用户登录参数的合法性
	memberDao := dao.MemberDao{tool.DBEngine}
	member := memberDao.Validate(loginParam)

	if member == nil {
		return nil
	}

	return member
}
