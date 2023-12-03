package views

import (
	"errors"
	"go_blog/common"
	"go_blog/service"
	"net/http"
	"strconv"
	"strings"
)

func (*HTMLApi) Detail(w http.ResponseWriter, r *http.Request) {
	detail := common.Template.Detail

	//获取路径参数
	path := r.URL.Path
	pIdStr := strings.TrimPrefix(path, "/p/")
	pIdStr = strings.TrimSuffix(pIdStr, ".html")
	pId, err := strconv.Atoi(pIdStr)
	if err != nil {
		detail.WriteError(w, errors.New("路径不匹配"))
		return
	}
	postRes, err := service.GetPostDetail(pId)
	if err != nil {
		detail.WriteError(w, errors.New("查询出错"))
		return
	}
	detail.WriteData(w, postRes)
}
