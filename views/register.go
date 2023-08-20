package views

import (
	"go-blog/common"
	"go-blog/config"
	"net/http"
)

func (*HTMLApi) Register(w http.ResponseWriter, r *http.Request) {
	print("1234")
	register := common.Template.Register
	register.WriteData(w, config.Cfg.Viewer)
}
