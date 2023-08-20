package views

import (
	"errors"
	"go-blog/common"
	"go-blog/service"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func (*HTMLApi) Detail(w http.ResponseWriter, r *http.Request) {
	detail := common.Template.Detail

	path := r.URL.Path
	pId, err := strconv.Atoi(strings.TrimSuffix(strings.TrimPrefix(path, "/p/"), ".html"))
	if err != nil {
		log.Println("路径错误")
		detail.WriteError(w, errors.New("URL格式错误"))
		return
	}
	dr, _ := service.GetPostDetail(pId)
	detail.WriteData(w, dr)
}
