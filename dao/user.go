package dao

import (
	"fmt"
	"go-blog/models"
	"log"
)

func GetUserNameById(userId int) string {
	row := DB.QueryRow("select user_name from blog_user where uid = ?", userId)
	if row.Err() != nil {
		log.Println(row.Err())
	}
	var userName string
	row.Scan(&userName)
	return userName
}

func GetUser(userName, passwd string) *models.User {
	row := DB.QueryRow(
		"select * from blog_user where user_name = ? and passwd=? limit 1",
		userName,
		passwd)
	if row.Err() != nil {
		log.Println(row.Err())
		return nil
	}
	var user models.User
	err := row.Scan(&user.Uid,
		&user.UserName,
		&user.Passwd,
		&user.Avatar,
		&user.CreateAt,
		&user.UpdateAt,
	)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	return &user
}

func CreateUser(user *models.User, passwdHash string) (*models.User, error) {
	fmt.Println("1234")
	result, err := DB.Exec(
		"insert into blog_user(user_name,passwd,avatar,create_at,update_at)"+
			"values(?,?,?,?,?) ",
		user.UserName,
		passwdHash,
		user.Avatar,
		user.CreateAt,
		user.UpdateAt,
	)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	uid, _ := result.LastInsertId()
	user.Uid = int(uid)
	return user, nil
}
