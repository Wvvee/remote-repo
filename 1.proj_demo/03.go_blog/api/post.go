package api

import (
	"errors"
	"go_blog/common"
	"go_blog/dao"
	"go_blog/models"
	"go_blog/service"
	"go_blog/utils"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func (*Api) GetPost(w http.ResponseWriter, r *http.Request) {

	//获取路径参数
	path := r.URL.Path
	pIdStr := strings.TrimPrefix(path, "/api/v1/post/")
	pId, err := strconv.Atoi(pIdStr)
	if err != nil {
		common.Error(w, errors.New("不识别此请求路径"))
		return
	}
	post, err := dao.GetPostByPid(pId)
	if err != nil {
		common.Error(w, err)
		return
	}
	common.Success(w, post)
}

//发布文章
func (*Api) SaveAndUpdatePost(w http.ResponseWriter, r *http.Request) {
	//获取用户id，判断用户是否登录,解析token，判断登录会话是否已经过期
	token := r.Header.Get("Authorization")
	_, claim, err := utils.ParseToken(token)
	if err != nil {
		common.Error(w, errors.New("登录已过期"))
		return
	}
	uid := claim.Uid
	//post操作 save

	method := r.Method
	//针对不同的请求做出不同的操作
	switch method {

	//post请求
	case http.MethodPost:
		params := common.GetRequestJsonParam(r)

		//因为不知道传过来cId里面具体interface{}装的什么类型，对其进行判断处理。
		cIdInterface := params["categoryId"]
		cId := common.GetType(cIdInterface)

		content := params["content"].(string)
		markdown := params["markdown"].(string)
		slug := params["slug"].(string)
		title := params["title"].(string)

		//对type进行类型判断
		postTypeInterface := params["type"]
		postType := common.GetType(postTypeInterface)

		post := &models.Post{
			Pid:        -1,
			Title:      title,
			Slug:       &slug,
			Content:    content,
			Markdown:   markdown,
			CategoryId: cId,
			UserId:     uid,
			ViewCount:  0,
			Type:       postType,
			CreateAt:   time.Now(),
			UpdateAt:   time.Now(),
		}
		service.SavePost(post)
		common.Success(w, post)

	//put请求--其实对应update
	case http.MethodPut:
		params := common.GetRequestJsonParam(r)

		//因为不知道传过来cId里面具体interface{}装的什么类型，对其进行判断处理。
		cIdInterface := params["categoryId"]
		cId := common.GetType(cIdInterface)

		content := params["content"].(string)
		markdown := params["markdown"].(string)
		slug := params["slug"].(string)
		title := params["title"].(string)

		//对type进行类型判断
		postTypeInterface := params["type"]
		postType := common.GetType(postTypeInterface)

		pidFloate := params["pid"].(float64)
		pid := int(pidFloate)
		post := &models.Post{
			Pid:        pid,
			Title:      title,
			Slug:       &slug,
			Content:    content,
			Markdown:   markdown,
			CategoryId: cId,
			UserId:     uid,
			ViewCount:  0,
			Type:       postType,
			CreateAt:   time.Now(),
			UpdateAt:   time.Now(),
		}
		service.UpdatePost(post)
		common.Success(w, post)
	}

}

func (*Api) SearchPost(w http.ResponseWriter, r *http.Request) {

	_ = r.ParseForm()
	condition := r.Form.Get("val")
	searchResp := service.SearchPost(condition)
	common.Success(w, searchResp)
}
