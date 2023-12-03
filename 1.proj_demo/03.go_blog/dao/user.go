package dao

import (
	"go_blog/models"
	"log"
)

func GetUserNameById(userId int) string {
	var userName string
	row := DB.QueryRow("select user_name from blog_user where uid = ?", userId)
	err := row.Scan(&userName)
	if err != nil {
		log.Println(row.Err())
	}
	return userName
}

func GetUser(userName string, passwd string) *models.User {
	row := DB.QueryRow("select * from blog_user where user_name = ? and passwd = ?", userName, passwd)
	if row.Err() != nil {
		log.Println("从数据库查询失败")
		return nil
	}
	var user = &models.User{}
	err := row.Scan(&user.Uid, &user.UserName, &user.Passwd, &user.Avatar, &user.CreateAt, &user.UpdateAt)
	if err != nil {
		log.Println("row.Scan()失败", err)
		return nil
	}
	return user
}
