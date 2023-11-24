package controllers

import (
	"gin-ranking/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func (UserController) GetUserInfo(c *gin.Context) {
	idStr := c.Param("id")
	name := c.Param("name")
	id, _ := strconv.Atoi(idStr)

	user, _ := models.GetUserTest(id)

	ReturnSuccess(c, 0, name, user, 1)
}

func (UserController) AddUser(c *gin.Context) {
	username := c.DefaultPostForm("username", "")
	id, err := models.AddUser(username)
	if err != nil {
		ReturnError(c, 4002, "保存错误")
		return
	}
	ReturnSuccess(c, 0, "保存成功", id, 1)
}

func (UserController) UpdateUser(c *gin.Context) {
	username := c.DefaultPostForm("username", "")
	idStr := c.DefaultPostForm("id", "")
	id, _ := strconv.Atoi(idStr)
	models.UpdateUser(id, username)
	ReturnSuccess(c, 0, "更新成功", true, 1)
}

func (UserController) DeleteUser(c *gin.Context) {
	idStr := c.DefaultPostForm("id", "")
	id, _ := strconv.Atoi(idStr)
	err := models.DeleteUser(id)
	if err != nil {
		ReturnError(c, 4002, "删除错误")
		return
	}
	ReturnSuccess(c, 0, "删除成功", true, 1)
}

func (UserController) GetUserListTest(c *gin.Context) {
	users, err := models.GetUserListTest()
	if err != nil {
		ReturnError(c, 4004, "没有相关数据")
		return
	}
	ReturnSuccess(c, 0, "获取成功", users, 1)
}

func (UserController) GetList(c *gin.Context) {
	// logger.Write("日志信息", "user")
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		fmt.Println("捕获异常", err)
	// 	}
	// }()

	// num1 := 1
	// num2 := 0
	// num3 := num1 / num2
	// ReturnError(c, 4004, num3)
}
