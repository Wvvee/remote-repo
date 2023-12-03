package router

import (
	"go_blog/api"
	"go_blog/views"
	"net/http"
)

func Router() {
	//1.返回页面.views 2.返回数据 3.返回静态资源

	//响应服务请求
	http.HandleFunc("/", views.HTML.Index)

	//代表你如果想要访问resource资源就走Handle
	//StripPrefix指将文件资源映射到文件服务器
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))

	//文章分类
	http.HandleFunc("/c/", views.HTML.Category)

	//登录login
	http.HandleFunc("/login", views.HTML.Login)

	//登录 api login
	http.HandleFunc("/api/v1/login", api.API.Login)

	//查看文章详情
	http.HandleFunc("/p/", views.HTML.Detail)

	//写作 writing
	http.HandleFunc("/writing", views.HTML.Writing)

	//发布文章
	http.HandleFunc("/api/v1/post", api.API.SaveAndUpdatePost)

	//编辑文章
	http.HandleFunc("/api/v1/post/", api.API.GetPost)

	//上传图片
	http.HandleFunc("/api/v1/qiniu/token", api.API.QiniuToken)

	//文章归档
	http.HandleFunc("/pageonhole", views.HTML.Pigeonhole)

	//首页搜索功能
	http.HandleFunc("/api/v1/post/search", api.API.SearchPost)
}
