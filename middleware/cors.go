package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

/**
 * @Author: yirufeng
 * @Date: 2021/1/16 5:03 下午
 * @Desc: 跨域中间件
 **/


//TODO:设置跨域中间件
func CorsMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cors.New(cors.Config{
			AllowOrigins:     []string{"https://localhost.com"},
			AllowMethods:     []string{"PUT", "PATCH"},
			AllowHeaders:     []string{"Origin"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			AllowOriginFunc: func(origin string) bool {
				return origin == "https://github.com"
			},
			MaxAge: 12 * time.Hour,
		})
	}
}
