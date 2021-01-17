package middleware

import (
	"gin_blog/config"
	"gin_blog/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
 * @Author: yirufeng
 * @Date: 2021/1/16 5:03 下午
 * @Desc: jwt验证中间件
 **/

//在上下文中存储的用户信息
type UserInfo struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
}

//校验token
func JwtMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//获取用户请求的头部
		token := ctx.Request.Header.Get("Authorization")

		claim, code := util.VerifyValidToken(token)
		if code != config.SUCCESS {
			ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
				"code":    code,
				"message": config.GetMessage(code),
			})
			return
		}

		//说明验证通过，设置上下文
		ctx.Set("user", UserInfo{
			Id:       claim.UserId,
			Username: claim.Username,
		})
		ctx.Next()
	}
}
