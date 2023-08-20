package api

import (
	"go-blog/common"
	"go-blog/service"
	"net/http"
)

func (*APIHander) Login(w http.ResponseWriter, r *http.Request) {
	params := common.GetRequestJsonParam(r)
	userName := params["username"].(string)
	passwd := params["passwd"].(string)
	loninRes, err := service.Login(userName, passwd)
	if err != nil {
		common.Error(w, err)
		return
	}
	common.Success(w, loninRes)
}
