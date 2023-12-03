package views

import (
	"go_blog/common"
	"go_blog/service"
	"net/http"
)

func (*HTMLApi) Pigeonhole(w http.ResponseWriter, r *http.Request) {
	pigeonhole := common.Template.Pigeonhole

	//查询对应数据
	pigeonholeRes := service.FindPostPigeonhole()
	pigeonhole.WriteData(w, pigeonholeRes)
}
