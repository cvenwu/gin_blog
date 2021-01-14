package model

import (
	"gin_blog/utils/errmsg"
	"gorm.io/gorm"
)

/**
 * @Author: y  irufeng
 * @Date: 2021/1/11 6:42 下午
 * @Desc: 文章对应的结构体
 **/

type Article struct {
	gorm.Model
	Category Category `gorm:"foreignkey:Cid"`
	Title    string   `gorm:"type:varchar(100);not null" json:"title"`
	Cid      int      `gorm:"type:int;not null" json:"cid"`   //指向分类对应的id，做一个逻辑关系的关联
	Desc     string   `gorm:"type:varchar(200);" json:"desc"` //描述
	Content  string   `gorm:"type:longtext" json:"content"`   //文章对应的内容
	Img      string   `gorm:"type:varchar(100);" json:"img"`  //文章对应的图片
}

//添加文章
//返回的是一个code,以此向前端返回不同的消息
func CreateArticle(data *Article) int {

	//或者使用下一行注释掉的代码，或者使用我们的钩子函数
	err := db.Create(&data).Error

	//如果执行失败
	if err != nil {
		return errmsg.ERROR
	}

	//如果成功
	return errmsg.SUCCESS
}

//分类的id，每一页的大小，当前在第几页
func GetCategoryArticles(id int, pageSize int, pageNum int) ([]Article, int) {
	var cateArtList []Article
	err := db.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where("cid=?", id).Find(&cateArtList).Error

	if err != nil {
		return nil, errmsg.ERROR_CATE_NOT_EXIST
	}

	return cateArtList, errmsg.SUCCESS

}

func GetArticleInfo(id int) (Article, int) {
	var art Article
	err := db.Debug().Preload("Category").Where("id = ?", id).First(&art).Error
	if err != nil {
		return art, errmsg.ERROR_ART_NOT_EXIST
	}

	return art, errmsg.SUCCESS
}

func GetArticles(pageSize int, pageNum int) ([]Article, int) {
	var articles []Article

	//去数据库中查找，需要传入一个pageSize，也就是一页有多少个，
	//第二个调用offset方法，传入一个偏移量，固定写法：(页码-1)*页面大小
	//TODO:这里我们使用gorm提供的预加载模型
	err := db.Debug().Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&articles).Error

	//如果错误不是记录没找到
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR
	}

	return articles, errmsg.SUCCESS
}

//编辑文章
//根据传入的id来查找分类，然后修改分类对应的信息
func EditArticle(id int, data *Article) int {
	var article Article
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img
	//更新user的model对应的表，而后传入id，最后进行修改
	err := db.Model(&article).Where("id = ?", id).Update(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//删除文章
func DeleteArticle(id int) int {
	var article Article
	//删除分类的方法
	//TODO:为什么不限查询出来赋值给user而是直接可以链式调用删除user
	err := db.Debug().Where("id = ?", id).Delete(&article).Error
	if err != nil { //说明有错
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
