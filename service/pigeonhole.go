package service

import (
	"go-blog/config"
	"go-blog/dao"
	"go-blog/models"
)

func FindPostPigeonhole() models.PigeonholeRes {
	categorys, _ := dao.GetALLCategory()
	posts, _ := dao.GetALLPost()
	pageonholeMap := make(map[string][]models.Post)
	for _, post := range posts {
		at := post.CreateAt
		month := at.Format("2006-01")
		pageonholeMap[month] = append(pageonholeMap[month], post)
	}
	return models.PigeonholeRes{
		Viewer:       config.Cfg.Viewer,
		SystemConfig: config.Cfg.System,
		Categorys:    categorys,
		Lines:        pageonholeMap,
	}
}
