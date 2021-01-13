package model

import (
	"fmt"
	"gin_blog/utils"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

/**
 * @Author: yirufeng
 * @Date: 2021/1/11 5:24 下午
 * @Desc:
 **/

var db *gorm.DB
var err error

func InitDB() {
	db, err = gorm.Open(utils.Db, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.DbUser, utils.DbPassWord, utils.DbHost, utils.DbPort, utils.DbName))

	if err != nil {
		fmt.Println("------------------------------初始化数据库出错，请稍后再试------------------------------", err.Error())
		return
	}

	//禁用默认表名的复数形式
	db.SingularTable(true)

	//自动迁移数据表
	db.AutoMigrate(&User{}, &Category{}, &Article{})

	//设置连接池中最大闲置连接数量
	db.DB().SetMaxIdleConns(10)

	//设置数据库的最大连接数量
	db.DB().SetMaxOpenConns(100)

	//设置连接的最大可复用时间 建议不要高于gin的默认timeout的时间
	db.DB().SetConnMaxLifetime(10 * time.Second)

}
