package apiv1

import (
	"gin_blog/model"
	"gin_blog/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

/**
 * @Author: yirufeng
 * @Date: 2021/1/16 3:32 下午
 * @Desc:
 **/

//定义全局变量，因为所有的业务逻辑函数都会使用到
var (
	code int
)

/*
功能描述：用户注册
请求携带数据格式：
{
	"username": "用户名",
	"password": "密码",
}
响应数据格式：
*/
func UserSignUp(c *gin.Context) {
	code = util.SUCCESS

	//1. 获取json中的username 以及 password
	var user model.User
	//忽略错误
	_ = c.ShouldBindJSON(&user)

	//2. 创建用户并返回code
	code = model.CreateUser(user)

	//3. 返回code以及对应的msg
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": util.GetMessage(code),
	})
}

//用户登录
func UserLogin(c *gin.Context) {
	code = util.SUCCESS

	var user model.User
	//1. 接收传入过来的参数
	//忽略错误
	_ = c.ShouldBindJSON(&user)

	//2. 判断用户名与密码是否匹配
	code = model.CheckUsernamePassword(user)
	if code != util.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"code":    code,
			"message": util.GetMessage(code),
		})
		return
	}
	//3. TODO: 生成JWT的token并返回
	var token string
	token, code = util.ReleaseToken(user)
	if code != util.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"code":    code,
			"message": util.GetMessage(code),
		})
		return
	}

	//4. 返回code以及对应的msg
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": util.GetMessage(code),
		"token":   token,
	})
}

//用户编辑
//获取用户的id然后只允许用户编辑密码
func UserEdit(c *gin.Context) {
	code = util.SUCCESS
	//获取传送过来的id
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		code = util.UserParamFormatInvalid
	}

	//获取Json提交的修改后的数据

	//根据id查询用户并修改原来数据的内容

	//返回修改成功

}

//删除用户
func UserDelete(c *gin.Context) {

}

//修改密码
func UserForgetPassword(c *gin.Context) {

}

//忘记密码
func UserResetPassword(c *gin.Context) {

}
