package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

/**
 * @Author: yirufeng
 * @Date: 2021/1/11 4:11 下午
 * @Desc:
 **/

var (
	AppMode  string
	HttpPort string
	JwtKey   string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string
)

//init函数是作为一个包初始化的一个函数的接口
func init() {

	//TODO: 这里是相对路径
	//file, err := ini.Load("../config/config.ini") ❌写法
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查配置文件路径:", err.Error())
		return
	}
	//加载服务器
	LoadServer(file)
	//加载数据库
	LoadData(file)
}

func LoadServer(file *ini.File) {

	//调用MustString主要是因为我们要设置默认值，也就是说我们如果通过Section下面的Key取不到值我们直接就给一个默认值debug
	AppMode = file.Section("Server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("Server").Key("HttpPort").MustString(":3000")
	JwtKey = file.Section("Server").Key("JwtKey").MustString("d@%@$GGE%$Y$%YRdgtj9340j5343q112#")
}

func LoadData(file *ini.File) {
	Db = file.Section("Database").Key("Db").MustString("mysql")
	DbHost = file.Section("Database").Key("DbHost").MustString("127.0.0.1")
	DbPort = file.Section("Database").Key("DbPort").MustString("3306")
	DbUser = file.Section("Database").Key("DbUser").MustString("root")
	DbPassWord = file.Section("Database").Key("DbPassWord").MustString("123456")
	DbName = file.Section("Database").Key("DbName").MustString("ginblog")
}
