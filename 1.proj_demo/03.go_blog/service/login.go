package service

import (
	"errors"
	"go_blog/dao"
	"go_blog/models"
	"go_blog/utils"
)

func Login(userName, passwd string) (*models.LoginRes, error) {
	// passwd = utils.Md5Crypt(passwd, "mszlu")
	// fmt.Println()
	user := dao.GetUser(userName, passwd)
	if user == nil {
		return nil, errors.New("账号密码不正确")
	}
	uid := user.Uid

	//jwt生成token令牌
	//生成token
	token, err := utils.Award(&uid)
	if err != nil {
		return nil, errors.New("token未生成")
	}
	var userInfo models.UserInfo
	userInfo.Uid = user.Uid
	userInfo.UserName = user.UserName
	userInfo.Avatar = user.Avatar

	var lr = &models.LoginRes{
		Token:    token,
		UserInfo: userInfo,
	}
	return lr, nil

}
