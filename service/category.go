package service

import (
	"go-blog/config"
	"go-blog/dao"
	"go-blog/models"
	"html/template"
	"log"
)

func GetPostByCategoryId(page, pageSize, cid int) (*models.CategoryResponse, error) {
	categorys, err := dao.GetALLCategory()
	if err != nil {
		return nil, err
	}
	categoryName := dao.GetCategoryNameById(cid)

	posts, err := dao.GetCategoryPage(page, pageSize, cid)
	if err != nil {
		return nil, err
	}
	var postMores []models.PostMore
	for _, post := range posts {
		categoryName := dao.GetCategoryNameById(post.CategoryId)
		log.Println(categoryName)
		userName := dao.GetUserNameById(post.UserId)
		log.Println(userName)
		content := []rune(post.Content)
		if len(content) > 100 {
			content = content[0:100]
		}
		postMore := models.PostMore{
			Pid:          post.Pid,
			Title:        post.Title,
			Slug:         post.Slug,
			Content:      template.HTML(content),
			CategoryId:   post.CategoryId,
			CategoryName: categoryName,
			UserId:       post.UserId,
			UserName:     userName,
			ViewCount:    post.ViewCount,
			Type:         post.Type,
			CreateAt:     models.DateDay(post.CreateAt),
			UpdateAt:     models.DateDay(post.UpdateAt),
		}
		postMores = append(postMores, postMore)
	}
	//11  10 2  10 1 9 1  21 3
	//  (11-1)/10 + 1 = 2
	total := dao.CountGetAllPostByCategoryId(cid)
	pagesCount := (total-1)/10 + 1
	var pages []int
	for i := 0; i < pagesCount; i++ {
		pages = append(pages, i+1)
	}

	var hr = &models.CategoryResponse{
		HomeResponse: &models.HomeResponse{
			Viewer:    config.Cfg.Viewer,
			Categorys: categorys,
			Posts:     postMores,
			Total:     total,
			Page:      page,
			Pages:     pages,
			PageEnd:   page != pagesCount,
		},
		CategoryName: categoryName,
	}
	return hr, nil
}
