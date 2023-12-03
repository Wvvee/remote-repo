package service

import (
	"go_blog/config"
	"go_blog/dao"
	"go_blog/models"
	"html/template"
)

func GetAllIndexInfo(slug string, page int, pageSize int) (*models.HomeResponse, error) {

	//获取文章分类
	categorys, err := dao.GetAllCategory()
	if err != nil {
		return nil, err
	}

	//获取posts
	var posts []models.Post
	var postMores []models.PostMore
	var total int
	//如果slug为空
	if slug == "" {
		posts, err = dao.GetPostPage(page, pageSize)
		total = dao.CountGetAllPost()
	} else {
		posts, err = dao.GetPostPageBySlug(slug, page, pageSize)
		total = dao.CountGetAllPostBySlug(slug)
	}

	//将posts中的数据填充到postMores
	for _, post := range posts {

		categoryName := dao.GetCategoryNameById(post.CategoryId)
		userName := dao.GetUserNameById(post.UserId)
		content := []rune(post.Content)
		if len(content) > 100 {
			content = content[0:100]
		}

		var postMore models.PostMore
		postMore = models.PostMore{
			Pid:          post.Pid,
			Title:        post.Title,
			Content:      template.HTML(content),
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
		postMores = append(postMores, postMore)
	}

	//获取文章总数并计算页数

	pagesCount := (total-1)/10 + 1
	var pages []int
	for i := 0; i < pagesCount; i++ {
		pages = append(pages, i+1)
	}

	//页面响应
	var hr = &models.HomeResponse{
		Viewer:    config.Cfg.Viewer,
		Categorys: categorys,
		Posts:     postMores,
		Total:     total,
		Page:      page,
		Pages:     pages,
		PageEnd:   page != pagesCount,
	}
	return hr, nil
}
