package service

import (
	"go-blog/config"
	"go-blog/dao"
	"go-blog/models"
	"html/template"
	"log"
)

func SearchPost(condition string) []models.SearchResp {
	posts, _ := dao.GetSearchPost(condition)
	var searchRes []models.SearchResp
	for _, post := range posts {
		searchRes = append(searchRes, models.SearchResp{
			Pid:   post.Pid,
			Title: post.Title,
		})
	}
	return searchRes
}

func GetPostDetail(pid int) (*models.PostRes, error) {
	postdetail, err := dao.GetPostById(pid)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	categoryName := dao.GetCategoryNameById(postdetail.CategoryId)
	userName := dao.GetUserNameById(postdetail.UserId)
	postMore := models.PostMore{
		Pid:          postdetail.Pid,
		Title:        postdetail.Title,
		Slug:         postdetail.Slug,
		Content:      template.HTML(postdetail.Content),
		CategoryId:   postdetail.CategoryId,
		CategoryName: categoryName,
		UserId:       postdetail.UserId,
		UserName:     userName,
		ViewCount:    postdetail.ViewCount,
		Type:         postdetail.Type,
		CreateAt:     models.DateDay(postdetail.CreateAt),
		UpdateAt:     models.DateDay(postdetail.UpdateAt),
	}
	postRes := models.PostRes{
		SystemConfig: config.Cfg.System,
		Viewer:       config.Cfg.Viewer,
		Article:      postMore,
	}
	return &postRes, nil
}

func Writing() (*models.WritingRes, error) {
	categorys, err := dao.GetALLCategory()

	if err != nil {
		log.Println(err)
		return nil, err
	}
	writingRes := models.WritingRes{
		Title:     config.Cfg.Viewer.Title,
		CdnURL:    config.Cfg.System.CdnURL,
		Categorys: categorys,
	}
	return &writingRes, nil
}

func SavePost(post *models.Post) {
	dao.SavePost(post)
}
func UpdatePost(post *models.Post) {
	dao.UpdatePost(post)
}
