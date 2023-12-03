package views

import (
	"errors"
	"go_blog/common"
	"go_blog/service"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func (*HTMLApi) Index(w http.ResponseWriter, r *http.Request) {
	index := common.Template.Index

	//获取表单
	if err := r.ParseForm(); err != nil {
		log.Println("表单获取失败", err)
		index.WriteError(w, errors.New("系统错误,请联系管理员"))
		return
	}
	pageStr := r.Form.Get("page")
	page := 1
	if pageStr != "" {
		page, _ = strconv.Atoi(pageStr)
	}

	//获取slug
	path := r.URL.Path
	slug := strings.TrimPrefix(path, "/")

	//每页显示的数量
	pageSize := 10
	hr, err := service.GetAllIndexInfo(slug, page, pageSize)

	//获取数据
	if err != nil {
		log.Println(err, "首页获取数据出错")
		index.WriteError(w, errors.New("系统错误，请联系管理员"))
	}
	//写入数据
	index.WriteData(w, hr)

}
