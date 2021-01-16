package util

import (
	"fmt"
	"gin_blog/config"
	"gin_blog/model"
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
}

//生成token
func ReleaseToken(u model.User) (string, int) {
	//设置token的过期时间
	expiredTime := time.Now().Add(time.Second * time.Duration(config.JwtDuration))

	var myClaim = MyClaim{
		Username: u.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiredTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Subject:   config.JwtSubject,
			Issuer:    config.JwtIssuer,
			Id:        fmt.Sprintf("%d", u.ID), //使用用户的id作为token的id
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaim).SignedString([]byte(config.JwtKey))
	if err != nil {
		return "", FAIL
	}

	return token, SUCCESS
}

//校验token是否有效
func VerifyValidToken(token string) (*MyClaim, int) {
	//说明token格式不对
	if len(token) <= 7 {
		return nil, TokenFormatInvalid
	}
	token = token[7:]
	var myClaim MyClaim
	tempToken, err := jwt.ParseWithClaims(token, &myClaim, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.JwtKey), nil
	})
	if err != nil {
		return nil, TokenInValid //TODO:这里可以根据错误再进行细分
	}

	return tempToken.Claims.(*MyClaim), SUCCESS
}
