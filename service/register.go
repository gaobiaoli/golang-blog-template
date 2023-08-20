package service

import (
	"errors"
	"go-blog/dao"
	"go-blog/models"
	"go-blog/utils"
	"log"
)

func Register(user *models.User) (*models.LoginRes, error) {
	passwdHash := utils.Md5Crypt(user.Passwd, "mszlu")
	user, err := dao.CreateUser(user, passwdHash)
	if err != nil {
		return nil, errors.New("账号密码不正确")
	}
	uid := user.Uid
	token, err := utils.Award(&uid)
	if err != nil {
		log.Println("Token Failed")
	}
	var userInfo models.UserInfo
	userInfo.Uid = user.Uid
	userInfo.UserName = user.UserName
	userInfo.Avatar = user.Avatar
	var lr = &models.LoginRes{
		Token:    token,
		UserInfo: userInfo,
	}
	return lr, nil
}
