package router

import (
	"go-blog/api"
	"go-blog/logger"
	"go-blog/views"
	"net/http"
)

func Router() {

	http.HandleFunc("/", logger.Logger(views.HTML.Index))
	// http.HandleFunc("/index", views.HTML.Index)
	http.HandleFunc("/c/", views.HTML.CategoryIndex)
	http.HandleFunc("/login", views.HTML.Login)
	http.HandleFunc("/register", views.HTML.Register)
	http.HandleFunc("/p/", logger.Logger(views.HTML.Detail))
	http.HandleFunc("/writing", views.HTML.Writing)
	http.HandleFunc("/pigeonhole", views.HTML.Pigeonhole)
	http.HandleFunc("/api/v1/login", api.API.Login)
	http.HandleFunc("/api/v1/register", api.API.Register)
	http.HandleFunc("/api/v1/post", api.API.SaveAndUpdate)
	http.HandleFunc("/api/v1/post/", api.API.GetPostAndDelete)
	http.HandleFunc("/api/v1/qiniu/token", api.API.QiniuToken)
	http.HandleFunc("/api/v1/post/search", api.API.SearchPost)
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))
}
