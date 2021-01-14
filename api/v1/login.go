package v1

import (
	"gin_blog/middleware"
	"gin_blog/model"
	"gin_blog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
 * @Author: yirufeng
 * @Date: 2021/1/14 10:37 下午
 * @Desc:
 **/

func Login(ctx *gin.Context) {
	var data model.User
	ctx.ShouldBindJSON(&data)

	var token string
	var code int
	code = model.CheckLogin(data.Username, data.Password)
	if code == errmsg.SUCCESS {
		//生成一个token
		token, code = middleware.SetToken(data.Username)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"token":   token,
	})
}
