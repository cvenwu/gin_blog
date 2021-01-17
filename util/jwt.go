package util

import (
	"fmt"
	"gin_blog/config"
	"github.com/dgrijalva/jwt-go"
	"time"
)

/**
 * @Author: yirufeng
 * @Date: 2021/1/16 9:57 下午
 * @Desc:
 **/

type MyClaim struct {
	jwt.StandardClaims
	Username string `json:"username"`
	UserId   uint   `json:"id"`
}

//生成token
func ReleaseToken(uid uint, username string) (string, int) {
	//设置token的过期时间
	expiredTime := time.Now().Add(time.Second * time.Duration(config.JwtDuration))

	var myClaim = MyClaim{
		UserId:   uid,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiredTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Subject:   config.JwtSubject,
			Issuer:    config.JwtIssuer,
			Id:        fmt.Sprintf("%d", uid), //使用用户的id作为token的id
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaim).SignedString([]byte(config.JwtKey))
	if err != nil {
		return "", config.FAIL
	}

	return token, config.SUCCESS
}

//校验token是否有效
func VerifyValidToken(token string) (*MyClaim, int) {
	//说明token格式不对
	if len(token) <= 7 {
		return nil, config.TokenFormatInvalid
	}
	token = token[7:]
	var myClaim MyClaim
	tempToken, err := jwt.ParseWithClaims(token, &myClaim, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.JwtKey), nil
	})
	if err != nil {
		return nil, config.TokenInValid //TODO:这里可以根据错误再进行细分
	}

	return tempToken.Claims.(*MyClaim), config.SUCCESS
}
