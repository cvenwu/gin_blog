package router

import (
	apiv1 "gin_blog/api/v1"
	"github.com/gin-gonic/gin"
)

/**
 * @Author: yirufeng
 * @Date: 2021/1/16 10:26 上午
 * @Desc:
 **/

//建立路由关系
func CollectRouter(r *gin.Engine) {
	apiGroup := r.Group("/api", nil)
	v1 := apiGroup.Group("/v1", nil)
	{
		//用户模块
		userV1Group := v1.Group("/users", nil)
		{
			//用户注册
			userV1Group.POST("/", apiv1.UserSignUp)
			//用户登录
			userV1Group.POST("/login", apiv1.UserLogin)
			//编辑用户信息
			userV1Group.PUT("/", apiv1.UserEdit)
			//删除用户
			userV1Group.DELETE("/:id", apiv1.UserDelete)
			//忘记密码
			//修改密码
		}

		//分类模块
		categoryV1Group := v1.Group("/categories", nil)
		{
			//查看所有分类
			categoryV1Group.GET("/", nil)
			//编辑某一个分类
			categoryV1Group.PUT("/:id", nil)
			//删除某一个分类
			categoryV1Group.DELETE("/:id", nil)
		}
		//文章模块
		articleV1Group := v1.Group("/articles", nil)
		{
			articleV1Group.GET("/", nil)
		}

	}

}
