package util

import (
	"gin_blog/config"
	"golang.org/x/crypto/bcrypt"
)

/**
 * @Author: yirufeng
 * @Date: 2021/1/16 4:37 下午
 * @Desc:
 **/

//根据传入的字符串，使用bcrypt生成加密后的密码
func GenerateHashedPassword(password string) (string, int) {
	pwdHashed, err := bcrypt.GenerateFromPassword([]byte(password), config.BcryptCost)
	if err != nil {
		return "", config.UserPasswordHashedFailed
	}

	//说明加密成功
	return string(pwdHashed), config.SUCCESS
}

//验证用户传入的密码是否正确
//如果返回不是nil说明密码不正确
func CheckPasswordIsValid(hashedPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
