package main

import (
	"go_blog/common"
	"go_blog/router"
	"log"
	"net/http"
)

func init() {
	//模板加载
	common.LoadTemplate()
}

func main() {
	//程序入口
	//启动一个http服务
	server := http.Server{
		Addr: "localhost:8080",
	}

	//路由
	router.Router()

	//监听端口并启动服务
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
