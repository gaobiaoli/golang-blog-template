package api

import (
	"go-blog/common"
	"go-blog/models"
	"go-blog/service"
	"net/http"
	"time"
)

func (*APIHander) Register(w http.ResponseWriter, r *http.Request) {
	params := common.GetRequestJsonParam(r)
	userName := params["username"].(string)
	passwd := params["passwd"].(string)
	user := &models.User{
		Uid:      0,
		UserName: userName,
		Passwd:   passwd,
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}
	registerRes, err := service.Register(user)
	if err != nil {
		common.Error(w, err)
		return
	}
	common.Success(w, registerRes)
}
