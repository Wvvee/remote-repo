package views

import (
	"go_blog/common"
	"go_blog/service"
	"net/http"
)

func (*HTMLApi) Writing(w http.ResponseWriter, r *http.Request) {
	// 加载writing template模板
	writing := common.Template.Writing
	wr := service.Writing()
	writing.WriteData(w, wr)
}
