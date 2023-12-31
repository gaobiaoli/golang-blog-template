package views

import (
	"go-blog/common"
	"go-blog/service"
	"net/http"
)

func (*HTMLApi) Writing(w http.ResponseWriter, r *http.Request) {
	writing := common.Template.Writing
	wr, _ := service.Writing()
	writing.WriteData(w, wr)
}
