package service

import (
	"go_blog/config"
	"go_blog/dao"
	"go_blog/models"
	"html/template"
	"log"
)

func SavePost(post *models.Post) {
	dao.SavePost(post)
}

func UpdatePost(post *models.Post) {
	dao.UpdatePost(post)
}

func GetPostDetail(Pid int) (*models.PostRes, error) {
	post, err := dao.GetPostByPid(Pid)
	if err != nil {
		return nil, err
	}
	var postMore models.PostMore
	categoryName := dao.GetCategoryNameById(post.CategoryId)
	userName := dao.GetUserNameById(post.UserId)
	postMore = models.PostMore{
		Pid:          post.Pid,
		Title:        post.Title,
		Content:      template.HTML(post.Content),
		CategoryId:   post.CategoryId,
		CategoryName: categoryName,
		UserId:       post.UserId,
		UserName:     userName,
		ViewCount:    post.ViewCount,
		Type:         post.Type,
		Slug:         post.Slug,
		CreateAt:     models.DateDay(post.CreateAt),
		UpdateAt:     models.DateDay(post.UpdateAt),
	}
	postMore.Pid = post.Pid
	var postRes = &models.PostRes{
		Viewer:       config.Cfg.Viewer,
		SystemConfig: config.Cfg.System,
		Article:      postMore,
	}
	return postRes, nil
}

func Writing() (wr models.WritingRes) {
	wr.Title = config.Cfg.Viewer.Title
	wr.CdnURL = config.Cfg.System.CdnURL
	category, err := dao.GetAllCategory()
	if err != nil {
		log.Println(err)
		return
	}
	wr.Categorys = category
	return
}

func SearchPost(condition string) []models.SearchResp {
	posts, _ := dao.GetPostSearch(condition)
	var searchResps []models.SearchResp
	for _, p := range posts {
		searchResps = append(searchResps, models.SearchResp{
			Pid:   p.Pid,
			Title: p.Title,
		})
	}
	return searchResps
}
