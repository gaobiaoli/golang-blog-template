package api

import (
	"errors"
	"fmt"
	"go-blog/common"
	"go-blog/dao"
	"go-blog/models"
	"go-blog/service"
	"go-blog/utils"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func (*APIHander) SearchPost(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	condition := r.Form.Get("val")
	searchResp := service.SearchPost(condition)
	common.Success(w, searchResp)
}

func (*APIHander) GetPostAndDelete(w http.ResponseWriter, r *http.Request) {
	detail := common.Template.Writing

	path := r.URL.Path
	pid, err := strconv.Atoi(strings.TrimPrefix(path, "/api/v1/post/"))
	if err != nil {
		log.Println("路径错误")
		detail.WriteError(w, errors.New("URL格式错误"))
		return
	}
	method := r.Method
	switch method {
	case http.MethodGet:
		post, _ := dao.GetPostById(pid)
		common.Success(w, post)
	case http.MethodDelete:
		post, _ := dao.DeletePost(pid)
		common.Success(w, post)
	}

}

func (*APIHander) SaveAndUpdate(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	_, claim, err := utils.ParseToken(token)
	if err != nil {
		common.Error(w, errors.New("登录已过期"))
		return
	}
	uid := claim.Uid

	method := r.Method
	switch method {
	case http.MethodPost:
		params := common.GetRequestJsonParam(r)
		categoryId := int(params["categoryId"].(float64))
		content := params["content"].(string)
		markdown := params["markdown"].(string)
		slug := params["slug"].(string)
		title := params["title"].(string)
		postType := params["type"].(float64)
		pType := int(postType)
		post := &models.Post{
			Pid:        -1,
			Title:      title,
			Slug:       slug,
			Content:    content,
			Markdown:   markdown,
			CategoryId: categoryId,
			UserId:     uid,
			ViewCount:  0,
			Type:       pType,
			CreateAt:   time.Now(),
			UpdateAt:   time.Now(),
		}
		service.SavePost(post)
		common.Success(w, post)
	case http.MethodPut:
		fmt.Println("*****************")
		params := common.GetRequestJsonParam(r)
		categoryId := int(params["categoryId"].(float64))
		fmt.Println(params["categoryId"])
		content := params["content"].(string)
		markdown := params["markdown"].(string)
		slug := params["slug"].(string)
		title := params["title"].(string)
		postType := params["type"].(float64)
		pType := int(postType)
		pidFloat := params["pid"].(float64)
		pid := int(pidFloat)
		post := &models.Post{
			Pid:        pid,
			Title:      title,
			Slug:       slug,
			Content:    content,
			Markdown:   markdown,
			CategoryId: categoryId,
			UserId:     uid,
			ViewCount:  0,
			Type:       pType,
			CreateAt:   time.Now(),
			UpdateAt:   time.Now(),
		}
		service.UpdatePost(post)
		common.Success(w, post)
	}

}
