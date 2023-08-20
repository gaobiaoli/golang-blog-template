package views

import (
	"go-blog/common"
	"go-blog/service"
	"net/http"
)

func (*HTMLApi) Pigeonhole(w http.ResponseWriter, r *http.Request) {
	Pigeonhole := common.Template.Pigeonhole
	pigeonholeRes := service.FindPostPigeonhole()
	Pigeonhole.WriteData(w, pigeonholeRes)
}
