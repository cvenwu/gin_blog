package model

import "gorm.io/gorm"

/**
 * @Author: yirufeng
 * @Date: 2021/1/11 6:42 下午
 * @Desc: 文章对应的结构体
 **/

type Article struct {
	gorm.Model
	Category Category
	Title    string `gorm:"type:varchar(100);not null" json:"title"`
	Cid      int    `gorm:"type:int;not null" json:"cid"`   //指向分类对应的id，做一个逻辑关系的关联
	Desc     string `gorm:"type:varchar(200);" json:"desc"` //描述
	Content  string `gorm:"type:longtext" json:"content"`   //文章对应的内容
	Img      string `gorm:"type:varchar(100);" json:"img"`  //文章对应的图片
}
