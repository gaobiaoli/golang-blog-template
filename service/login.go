package service

import (
	"errors"
	"go-blog/dao"
	"go-blog/models"
	"go-blog/utils"
	"log"
)

func Login(username, passwd string) (*models.LoginRes, error) {
	passwd = utils.Md5Crypt(passwd, "mszlu")
	log.Println(passwd)
	user := dao.GetUser(username, passwd)
	if user == nil {
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
