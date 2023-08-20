package views

import (
	"errors"
	"go-blog/common"
	"go-blog/service"
	"log"
	"net/http"
	"strconv"
)

func (*HTMLApi) Index(w http.ResponseWriter, r *http.Request) {
	index := common.Template.Index
	if err := r.ParseForm(); err != nil {
		log.Println("表单获取失败")
		index.WriteError(w, errors.New("系统错误，请联系管理员"))
	}
	pageStr := r.Form.Get("page")
	page := 1
	if pageStr != "" {
		page, _ = strconv.Atoi(pageStr)
	}
	pageSize := 10
	hr, err := service.GetALLIndexInfo(page, pageSize)
	if err != nil {
		log.Println("Index Init Failer", err)
		index.WriteError(w, errors.New("系统错误，请联系管理员"))
	}
	index.WriteData(w, hr)
}
