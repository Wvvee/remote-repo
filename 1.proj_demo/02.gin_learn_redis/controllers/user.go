package controllers

import (
	"gin-ranking/models"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type UserController struct{}

type UserApi struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

func (UserController) Register(c *gin.Context) {
	//接收 用户名，密码，确认密码
	username := c.DefaultPostForm("username", "")
	password := c.DefaultPostForm("password", "")
	confirmPassword := c.DefaultPostForm("confirmPassword", "")

	if username == "" || password == "" || confirmPassword == "" {
		ReturnError(c, 4001, "请输入正确的信息")
		return
	}

	if password != confirmPassword {
		ReturnError(c, 4001, "密码和确认密码不一致")
		return
	}

	user, err := models.GetUserInfoByUsername(username)
	if user.Id != 0 {
		ReturnError(c, 4001, "username already exist.")
		return
	}

	_, err = models.AddUser(username, EncryMd5(password))
	if err != nil {
		ReturnError(c, 4001, "注册失败,请联系管理员")
		return
	}

	ReturnSuccess(c, 0, "注册成功", "", 1)
}

func (UserController) Login(c *gin.Context) {
	//receive username password
	username := c.DefaultPostForm("username", "")
	password := c.DefaultPostForm("password", "")

	if username == "" || password == "" {
		ReturnError(c, 4001, "请输入正确的信息")
		return
	}

	user, _ := models.GetUserInfoByUsername(username)
	if user.Id == 0 {
		ReturnError(c, 4004, "用户名或密码不正确")
		return
	}
	if user.Password != EncryMd5(password) {
		ReturnError(c, 4004, "用户名和密码不正确")
		return
	}

	session := sessions.Default(c)
	session.Set("login:"+strconv.Itoa(user.Id), user.Id)
	session.Save()
	data := UserApi{Id: user.Id, Username: user.Username}
	ReturnSuccess(c, 0, "恭喜，登陆成功", data, 1)
}
