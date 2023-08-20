package views

import (
	"errors"
	"go-blog/common"
	"go-blog/models"
	"go-blog/service"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func (*HTMLApi) CategoryIndex(w http.ResponseWriter, r *http.Request) {
	categoryindex := common.Template.Category
	categoryindex.WriteData(w, models.CategoryResponse{})
	if err := r.ParseForm(); err != nil {
		log.Println("表单获取失败")
		categoryindex.WriteError(w, errors.New("系统错误，请联系管理员"))
	}

	path := r.URL.Path
	categoryId, err := strconv.Atoi(strings.TrimPrefix(path, "/c/"))

	if err != nil {
		log.Println("路径错误")
		categoryindex.WriteError(w, errors.New("URL格式错误"))
	}

	pageStr := r.Form.Get("page")
	page := 1
	if pageStr != "" {
		page, _ = strconv.Atoi(pageStr)
	}
	pageSize := 10
	cr, err := service.GetPostByCategoryId(page, pageSize, categoryId)
	if err != nil {
		log.Println("Index Init Failer", err)
		categoryindex.WriteError(w, errors.New("系统错误，请联系管理员"))
	}
	categoryindex.WriteData(w, cr)
}
