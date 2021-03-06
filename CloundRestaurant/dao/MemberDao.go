package dao

import (
	"CloundRestaurant/tool"
	"CloundRestaurant/param"
	"CloundRestaurant/model"
	"log"
	"fmt"
)

type MemberDao struct {
	*tool.Orm
}

//上传用户头像
func (memberDao *MemberDao) UploadMemberAvator(userId int64, fileName string) int64 {

	member := model.Member{Avatar:fileName}

	result, err := memberDao.Where("id = ?", userId).Update(&member)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}

	return result
}

//验证用户是否存在
func (memberDao *MemberDao) Exsit(username string) bool {

	var member model.Member

	if _, err := memberDao.Where("username = ?", username).Get(&member); err != nil {
		return false
	}

	return true
}

//用户添加
func (memberDao *MemberDao) AddMember(member model.Member) int64 {

	result, err := memberDao.InsertOne(member)
	if err != nil {
		log.Fatal(err)
		return 0
	}

	return result
}

//验证用户合法性
func (memberDao *MemberDao) Validate(loginParam param.LoginParam) *model.Member {

	var member model.Member

	//获取用户名和密码
	username := loginParam.Username
	passwordHash := tool.Sha256(loginParam.Password)

	//判断合法性
	if _, err := memberDao.Where("username = ? and password = ?", username, passwordHash).Get(&member); err != nil {
		log.Fatal(err)
	}

	return &member
}
