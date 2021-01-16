package model

import (
	"gin_blog/common"
	"gin_blog/util"
	"gorm.io/gorm"
)

/**
 * @Author: yirufeng
 * @Date: 2021/1/16 3:30 下午
 * @Desc:
 **/

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}

//根据传入的username查询用户名是否存在
func UsernameIsExist(username string) bool {
	var user User
	db := common.GetDB()
	db.Where("username = ?", username).First(&user)

	//说明没找到用户
	if user.ID == 0 {
		return false
	}

	//说明用户名为username的用户是存在的
	return true
}

//添加传入的用户：如果用户名存在就不创建并返回用户名已存在的状态码，如果不存在就创建，创建成功返回success，如果密码加密失败返回加密失败
func CreateUser(user User) int {
	if UsernameIsExist(user.Username) {
		return util.UserNameExist
	}
	db := common.GetDB()

	//对密码进行加密
	passwordHashed, code := util.GenerateHashedPassword(user.Password)
	if code != util.SUCCESS {
		return util.UserPasswordHashedFailed
	}

	//然后保存加密后的密码
	user.Password = passwordHashed
	db.Create(&user)
	return util.SUCCESS
}

//用户登录验证
func CheckUsernamePassword(user User) int {
	//如果用户名不存在
	if !UsernameIsExist(user.Username) {
		return util.UserNameNotExist
	}
	//查询对应用户名
	var u User
	db := common.GetDB()
	db.Where("username = ?", user.Username).First(&u)
	if util.CheckPasswordIsValid(u.Password, user.Password) != nil {
		return util.UserPasswordNotRight
	}
	return util.SUCCESS
}
