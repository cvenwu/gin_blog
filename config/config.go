package config

import (
	"fmt"
	"gopkg.in/ini.v1"
)

/**
 * @Author: yirufeng
 * @Date: 2021/1/16 3:49 下午
 * @Desc:
 **/

var (
	//服务器相关
	ServerHost string //部署到的服务器的ip
	ServerPort string //部署到的服务器的端口

	//用户密码相关
	BcryptCost int //用户密码加密需要的盐

	//数据库相关
	DbType     string //数据库种类
	DbHost     string //数据库对应ip地址
	DbPort     string //数据库对应端口
	DbUsername string //数据库用户名
	DbPassword string //数据库密码
	DbName     string //数据库名称

	//JWT相关
	JwtKey      string //jwt加密所使用的key
	JwtDuration int
	JwtSubject  string
	JwtIssuer   string
)

func InitConfig() {
	f, err := ini.Load("/config/config.ini")
	if err != nil {
		fmt.Println("-----------------------ini配置文件读取出错，请稍后再试-----------------------", err.Error())
		return
	}

	//如果不存在就返回零值
	ServerHost = f.Section("server").Key("Host").MustString("127.0.0.1")
	ServerPort = f.Section("server").Key("Port").MustString(":8080")

	DbType = f.Section("mysql").Key("Name").MustString("mysql")
	DbHost = f.Section("mysql").Key("Host").MustString("127.0.0.1")
	DbPort = f.Section("mysql").Key("Port").MustString("3306")
	DbUsername = f.Section("mysql").Key("User").MustString("root")
	DbPassword = f.Section("mysql").Key("Password").MustString("123456")
	DbName = f.Section("mysql").Key("Name").MustString("gin_vue_blog")

	JwtKey = f.Section("jwt").Key("JwtKey").MustString("e0e879b62d2b43c353bf9ac50a057ec928e51133752c9b500eaa95e4986a57fd")
	JwtDuration = f.Section("jwt").Key("Duration").MustInt(24 * 3600)
	JwtSubject = f.Section("jwt").Key("Subject").MustString("blog_vue")
	JwtIssuer = f.Section("jwt").Key("Issuer").MustString("admin")

	BcryptCost = f.Section("user-password").Key("BcryptCost").MustInt(10) //默认是10
}
