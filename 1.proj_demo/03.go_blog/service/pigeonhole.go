package service

import (
	"go_blog/config"
	"go_blog/dao"
	"go_blog/models"
)

func FindPostPigeonhole() models.PigeonholeRes {
	// 查询所有文章，进行月份整理
	//查询所有分类

	posts, _ := dao.GetPostAll()
	pigeonholeMap := make(map[string][]models.Post)
	for _, post := range posts {
		at := post.CreateAt
		month := at.Format("2006-01")
		pigeonholeMap[month] = append(pigeonholeMap[month], post)
	}

	categorys, _ := dao.GetAllCategory()
	return models.PigeonholeRes{
		Viewer:       config.Cfg.Viewer,
		SystemConfig: config.Cfg.System,
		Categorys:    categorys,
		Lines:        pigeonholeMap,
	}
}
