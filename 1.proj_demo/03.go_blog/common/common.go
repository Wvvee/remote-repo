package common

import (
	"encoding/json"
	"go_blog/config"
	"go_blog/models"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var Template models.HtmlTemplate

//加载vue和js模板
func LoadTemplate() {
	w := sync.WaitGroup{}
	w.Add(1)
	go func() {
		//将耗时操作加入协程
		var err error
		Template, err = models.InitTemplate(config.Cfg.System.CurrentDir + "/template/")
		if err != nil {
			panic(err)
		}
		w.Done()
	}()
	w.Wait()
}

func Error(w http.ResponseWriter, err error) {
	var result models.Result
	result.Code = -999
	result.Error = err.Error()
	resultJson, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(resultJson)
	if err != nil {
		log.Println(err)
	}
}

//创建一个成功的HTTP响应，其中包含了状态码200，没有错误，以及传入的数据
func Success(w http.ResponseWriter, data interface{}) {
	var result models.Result
	result.Code = 200
	result.Error = ""
	result.Data = data
	resultJson, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write(resultJson)
	if err != nil {
		log.Println(err)
	}
}

func GetRequestJsonParam(r *http.Request) map[string]interface{} {
	var params map[string]interface{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err, "body 解析出现错误")
	}

	// fmt.Println("原始请求体：", string(body))

	err = json.Unmarshal(body, &params)
	if err != nil {
		log.Println(err, "JSON 解析出现错误")
	}
	return params
}

//判断接口存放值类型
func GetType(postTypeInterface interface{}) int {
	var postType int
	switch v := postTypeInterface.(type) {
	case int:
		postType = v
	case string:
		postType, _ = strconv.Atoi(v)
	case float64:
		postType = int(v)
	default:
		postType = 0
	}

	return postType
}
