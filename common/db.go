package common

import (
	"fmt"
	"gin_blog/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

/**
 * @Author: yirufeng
 * @Date: 2021/1/16 4:03 下午
 * @Desc:
 **/

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(config.DbType,
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			config.DbUsername, config.DbPassword,
			config.DbHost, config.DbPort,
			config.DbName))
	if err != nil {
		fmt.Println("-----------------------数据库初始化失败，请稍后再试-----------------------", err.Error())
		return
	}
}

func GetDB() *gorm.DB {
	return DB
}
