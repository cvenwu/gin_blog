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
		//添加用户
		router.POST("user/add", v1.AddUser)
		//查询用户列表
		router.GET("users", v1.GetUsers)
		//编辑用户
		router.PUT("user/:id", v1.EditUser)
		//删除用户
		router.DELETE("user/:id", v1.DeleteUser)

		//category模块的路由接口
		//添加分类
		router.POST("category/add", v1.AddCategory)
		//查询分类列表
		router.GET("categories", v1.GetCategories)
		//编辑分类
		router.PUT("category/:id", v1.EditCategory)
		//删除分类
		router.DELETE("category/:id", v1.DeleteCategory)

		//article模块的路由接口
		//添加文章
		router.POST("article/add", v1.AddArticle)
		//查询文章列表
		router.GET("articles", v1.GetArticles)
		//编辑文章
		router.PUT("article/:id", v1.EditArticle)
		//查询单个文章
		router.GET("article/info/:id", v1.GetArticle)

		//删除文章
		router.DELETE("article/:id", v1.DeleteArticle)
		//查询某一个分类下的所有文章
		router.GET("article/list/:id", v1.GetCategoryArticles)
	}

	r.Run(utils.HttpPort)
}
