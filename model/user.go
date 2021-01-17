package model

import (
	"gin_blog/common"
	"gin_blog/config"
	"gin_blog/middleware"
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
		return config.UserNameExist
	}
	db := common.GetDB()

	//对密码进行加密
	passwordHashed, code := util.GenerateHashedPassword(user.Password)
	if code != config.SUCCESS {
		return config.UserPasswordHashedFailed
	}

	//然后保存加密后的密码
	user.Password = passwordHashed
	db.Create(&user)
	return config.SUCCESS
}

//用户登录验证
func CheckUsernamePassword(user User) int {
	//如果用户名不存在
	if !UsernameIsExist(user.Username) {
		return config.UserNameNotExist
	}
	//查询对应用户名
	var u User
	db := common.GetDB()
	db.Where("username = ?", user.Username).First(&u)
	if util.CheckPasswordIsValid(u.Password, user.Password) != nil {
		return config.UserPasswordNotRight
	}
	return config.SUCCESS
}

//修改用户信息：因为密码我们单独做了一个功能，所以目前这里只支持修改用户名
func EditUser(u middleware.UserInfo) int {
	db := common.GetDB()
	//TODO:这里不确定是否写对
	ret := db.Debug().Where("id = ?", u.Id).Select("username").Update("username", u.Username).RowsAffected
	if ret <= 0 {
		return config.FAIL
	}
	return config.SUCCESS
}

//删除用户
func DeleteUser(id int) int {
	if id <= 0 { //id非法
		return config.UserParamFormatInvalid
	}
	var u User
	db := common.GetDB()
	//根据id删除用户
	//TODO:这里是否可以优化
	affected := db.Debug().Where("id = ?", id).First(&u).Delete(&u).RowsAffected
	if affected <= 0 {
		return config.FAIL
	}
	return config.SUCCESS
}
