package api

import (
	"go_blog/common"
	"go_blog/service"
	"net/http"
)

func (*Api) Login(w http.ResponseWriter, r *http.Request) {
	// 接收登录信息的用户名密码和返回对应json数据
	params := common.GetRequestJsonParam(r)
	// fmt.Println(params)
	var userName string
	var passwd string
	userName, _ = params["username"].(string)
	passwd, _ = params["passwd"].(string)
	// fmt.Println(userName, passwd)
	loginRes, err := service.Login(userName, passwd)
	if err != nil {
		common.Error(w, err)
	}
	common.Success(w, loginRes)
}
