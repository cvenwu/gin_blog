package model

import (
	"encoding/base64"
	"fmt"
	"gin_blog/utils/errmsg"
	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
	"log"
)

/**
 * @Author: yirufeng
 * @Date: 2021/1/11 5:08 下午
 * @Desc:
 **/

type User struct {
	gorm.Model
	//以下3个tag写上json主要是为了前后端数据交互使用
	Username string `gorm:"type:varchar(20);not null" json:"username"`
	Password string `gorm:"type:varchar(20);not null" json:"password"`
	Role     int    `gorm:"type:int" json:"role"`
}

//查询用户是否存在
func CheckUserExist(username string) int {
	var user User
	//根据username查询出第一个符合用户名的用户并赋值给user，会将select查询的字段赋值给user
	//也就是说只会查询Select参数中指定的字段，然后给user进行赋值
	db.Debug().Select("id").Where("username = ?", username).First(&user)

	//看一下是否只是查询id，还是都会查询
	fmt.Println(user)
	if user.ID > 0 { //说明用户名不可以使用
		return errmsg.ERROR_USERNAME_USED
	}

	//用户名不存在
	return errmsg.SUCCESS
}

//添加用户
//返回的是一个code,以此向前端返回不同的消息
func CreateUser(data *User) int {

	//或者使用下一行注释掉的代码，或者使用我们的钩子函数
	data.Password = ScryptPassword(data.Password)
	err := db.Create(&data).Error

	//如果执行失败
	if err != nil {
		return errmsg.ERROR
	}

	//如果成功
	return errmsg.SUCCESS
}

//查询用户列表，需要分页，根据前端query传入一个pageSize，以及当前的页码pageNum
func GetUsers(pageSize int, pageNum int) []User {
	var users []User

	//去数据库中查找，需要传入一个pageSize，也就是一页有多少个，
	//第二个调用offset方法，传入一个偏移量，固定写法：(页码-1)*页面大小
	err = db.Debug().Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error

	//如果错误不是记录没找到
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}

	return users
}

//编辑用户
//根据传入的id来查找用户，然后修改用户对应的信息
func EditUser(id int, data *User) int {
	var user User
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role
	//更新user的model对应的表，而后传入id，最后进行修改
	err = db.Model(&user).Where("id = ?", id).Update(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//删除用户
func DeleteUser(id int) int {
	var user User
	//删除用户的方法
	//TODO:为什么不限查询出来赋值给user而是直接可以链式调用删除user
	err = db.Debug().Where("id = ?", id).Delete(&user).Error
	if err != nil { //说明有错
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//钩子函数：对密码进行加密
//func (u *User) BeforeSave() {
//	u.Password = ScryptPassword(u.Password)
//}

//对密码加密
func ScryptPassword(password string) string {
	//最后的加密之后的位数
	const KeyLen = 10
	//生成盐值
	salt := []byte{12, 32, 4, 6, 6, 22, 222, 11}

	HashPwd, err := scrypt.Key([]byte(password), salt, 1<<14, 8, 1, KeyLen)

	if err != nil {
		log.Fatal(err)
	}

	//参考使用说明的demo
	return base64.StdEncoding.EncodeToString(HashPwd)
}
