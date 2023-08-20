package api

import (
	"go-blog/common"
	"go-blog/config"
	"net/http"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

func (*APIHander) QiniuToken(w http.ResponseWriter, r *http.Request) {
	mac := qbox.NewMac(config.Cfg.System.QiniuAccessKey, config.Cfg.System.QiniuSecretKey)
	putPolicy := storage.PutPolicy{
		Scope: config.Cfg.System.Bucket,
	}
	upToken := putPolicy.UploadToken(mac)
	common.Success(w, upToken)
}
