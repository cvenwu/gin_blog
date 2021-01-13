package model

import "gorm.io/gorm"

/**
 * @Author: yirufeng
 * @Date: 2021/1/11 6:40 下午
 * @Desc:
 **/


//设置文章对应的分五类
type Category struct {
	gorm.Model
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}