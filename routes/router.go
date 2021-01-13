package routes

import (
	v1 "gin_blog/api/v1"
	"gin_blog/utils"
	"github.com/gin-gonic/gin"
)

/**
 * @Author: yirufeng
 * @Date: 2021/1/11 4:47 下午
 * @Desc: 编写路由的入口文件
 **/

func InitRouter() {
	//设置gin的模式
	gin.SetMode(utils.AppMode)

	//初始化路由
	r := gin.Default()

	//因为是前后端分离，所以这里需要一个版本，所以需要一个路由组来做
	router := r.Group( "api/v1")
	{
		//此时下面的路由前缀都必须是api/v1再加上我们这里的相对路由才可以访问的到

		//user模块的路由接口
		router.POST("user/add", v1.AddUser)
		//查询用户列表
		router.GET("users", v1.GetUsers)
		//编辑用户
		router.PUT("user/:id", v1.EditUser)
		//删除用户
		router.DELETE("user/:id", v1.DeleteUser)

		//category模块的路由接口

		//article模块的路由接口
	}

	r.Run(utils.HttpPort)
}
