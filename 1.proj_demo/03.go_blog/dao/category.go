package dao

import (
	"go_blog/models"
	"log"
)

func GetCategoryNameById(cId int) string {
	var categoryName string
	row := DB.QueryRow("select name from blog_category where cid = ?", cId)
	err := row.Scan(&categoryName)
	if err != nil {
		log.Println(row.Err())
	}
	return categoryName
}

//获取所有文章种类
func GetAllCategory() ([]models.Category, error) {
	rows, err := DB.Query("select * from blog_category")
	if err != nil {
		log.Println("GetAllCategory 查询出错", err)
		return nil, err
	}
	var categorys []models.Category
	for rows.Next() {
		var category models.Category
		err = rows.Scan(&category.Cid, &category.Name, &category.CreateAt, &category.UpdateAt)
		if err != nil {
			log.Println("GetAllCategory 取值出错", err)
			return nil, err
		}
		categorys = append(categorys, category)
	}
	return categorys, err
}
