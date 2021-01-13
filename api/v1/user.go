package v1

import (
	"gin_blog/model"
	"gin_blog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

/**
 * @Author: yirufeng
 * @Date: 2021/1/13 3:59 下午
 * @Desc:
 **/

//声明一个全局的code
var code int

//添加用户
func AddUser(c *gin.Context) {
	//1. 拿到用户
	var user model.User

	//绑定模型
	//如果不需要冒黄我们就不需要接收
	_ = c.ShouldBindJSON(&user)

	//2. 判断用户名是否存在
	code = model.CheckUserExist(user.Username)
	if code == errmsg.SUCCESS { //如果不存在就加入
		model.CreateUser(&user) // 写入数据库
	}
	//3.如果已经被使用了
	if code == errmsg.ERROR_USERNAME_USED {
		code = errmsg.ERROR_USERNAME_USED
	}

	//第一个200是网络传输的200，与我们业务操作状态码无关
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    user,
		"message": errmsg.GetErrMsg(code),
	})
}

//查询单个用户：查询某一个用户的所有信息
//这里先不写，因为意义不大

//查询用户列表
func GetUsers(c *gin.Context) {
	//接收前端传过来的两个query
	pageSize, _ := strconv.Atoi(c.Query("pageSize")) //返回的是一个string
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))   //返回的是一个string

	//gorm官网：gorm.io/docs/query.html中有说到如果传入-1将会取消limit限制
	if pageSize == 0 {
		pageSize = -1
	}

	if pageNum == 0 {
		pageNum = -1
	}

	data := model.GetUsers(pageSize, pageNum)
	code = errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

//编辑用户：用户传入的信息一定是想要修改后的信息，
func EditUser(c *gin.Context) {

	//获取用户名是否被占用
	var data model.User
	c.ShouldBindJSON(&data)
	id, _ := strconv.Atoi(c.Param("id"))


	//TODO:这里也可以使用钩子函数改进，也就是在save之前check user
	code = model.CheckUserExist(data.Username)
	if code == errmsg.SUCCESS {
		model.EditUser(id, &data)
	}
	if code == errmsg.ERROR_USERNAME_USED {
		c.Abort()
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

//删除用户
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteUser(id)

	//可以直接返回,因为只需要一个成功或失败
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
