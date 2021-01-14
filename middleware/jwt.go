package middleware

import (
	"gin_blog/utils"
	"gin_blog/utils/errmsg"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

/**
 * @Author: yirufeng
 * @Date: 2021/1/14 8:27 下午
 * @Desc:
 **/

//秘钥的参数
var JwtKey = []byte(utils.JwtKey)

//接收的参数
type MyClaims struct {
	//需要与我们的用户模型定义的字段保持一致
	Username string `json:"username"`
	jwt.StandardClaims
}

var code int

//生成token
//第一个返回的是token，第二个是错误码
func SetToken(username string) (string, int) {
	expireTime := time.Now().Add(10 * time.Hour)

	//新建一个结构体
	setClaims := MyClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "yirufeng", //签发人
		},
	}

	reqClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, setClaims)
	token, err := reqClaims.SignedString(JwtKey)
	if err != nil {
		return "", errmsg.ERROR
	}
	return token, errmsg.SUCCESS
}

//验证token
//也提供了一个方法jwt.ParseWithClaims()
func CheckToken(token string) (*MyClaims, int) {

	//拿到的是一个token结构体
	settoken, _ := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})

	if key, _ := settoken.Claims.(*MyClaims); settoken.Valid {
		return key, errmsg.SUCCESS
	}

	return nil, errmsg.ERROR

}

//jwt中间件
//token有固定格式
func JwtToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenHeader := ctx.Request.Header.Get("Authorization") //是一个规范

		code = errmsg.SUCCESS

		//返回一个code
		if tokenHeader == "" {
			code = errmsg.ERROR_TOKEN_EXIST
			ctx.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			ctx.Abort()
			return
		}

		checkToken := strings.SplitN(tokenHeader, " ", 2)
		if len(checkToken) != 2 && checkToken[0] != "Bearer" {
			code = errmsg.ERROR_TOKEN_TYPE_WRONG
			ctx.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			ctx.Abort()
			return
		}

		claim, ok := CheckToken(checkToken[1])
		if ok == errmsg.ERROR {
			code = errmsg.ERROR_TOKEN_WRONG
			ctx.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			ctx.Abort()
			return
		}
		if time.Now().Unix() > claim.ExpiresAt {
			code = errmsg.ERROR_TOKEN_RUNTIME
			ctx.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			ctx.Abort()
			return
		}



		//设置上下文
		ctx.Set("username", claim.Username)
		ctx.Next()
	}
}
