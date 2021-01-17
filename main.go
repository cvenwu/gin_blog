package main

import (
	"gin_blog/common"
	"gin_blog/config"
	"gin_blog/middleware"
	"gin_blog/router"
	"github.com/gin-gonic/gin"
)

/**
 * @Author: yirufeng
 * @Date: 2021/1/16 10:25 上午
 * @Desc:
 **/

func main() {
	//初始化配置
	config.InitConfig()
	//初始化数据库
	common.InitDB()
	//配置运行方式
	gin.SetMode(config.RunMode)
	r := gin.Default()
	//使用中间件
	r.Use(middleware.JwtMiddleWare(), middleware.CorsMiddleWare())
	//使用路由
	router.CollectRouter(r)
	r.Run(config.ServerHost + config.ServerPort)
}
