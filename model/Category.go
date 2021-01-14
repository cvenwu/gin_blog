package model

import (
	"gin_blog/utils/errmsg"
	"github.com/jinzhu/gorm"
)

/**
 * @Author: yirufeng
 * @Date: 2021/1/11 6:40 下午
 * @Desc:
 **/

//设置文章对应的分五类
type Category struct {
	ID   uint   `gorm:"primaryKey;auto_increment"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

//查询分类是否存在
func CheckCategory(name string) int {
	var cate Category
	//根据username查询出第一个符合用户名的用户并赋值给user，会将select查询的字段赋值给user
	//也就是说只会查询Select参数中指定的字段，然后给user进行赋值
	db.Debug().Select("id").Where("name = ?", name).First(&cate)

	//看一下是否只是查询id，还是都会查询
	if cate.ID > 0 { //说明用户名不可以使用
		return errmsg.ERROR_CATEGORY_USED
	}

	//用户名不存在
	return errmsg.SUCCESS
}

//添加分类
//返回的是一个code,以此向前端返回不同的消息
func CreateCategory(data *Category) int {

	//或者使用下一行注释掉的代码，或者使用我们的钩子函数
	err := db.Create(&data).Error

	//如果执行失败
	if err != nil {
		return errmsg.ERROR
	}

	//如果成功
	return errmsg.SUCCESS
}

//查询分类列表，需要分页，根据前端query传入一个pageSize，以及当前的页码pageNum
func GetCategories(pageSize int, pageNum int) []Category {
	var cates []Category

	//去数据库中查找，需要传入一个pageSize，也就是一页有多少个，
	//第二个调用offset方法，传入一个偏移量，固定写法：(页码-1)*页面大小
	err = db.Debug().Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&cates).Error

	//如果错误不是记录没找到
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}

	return cates
}


//编辑分类
//根据传入的id来查找分类，然后修改分类对应的信息
func EditCategory(id int, data *Category) int {
	var cate Category
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	//更新user的model对应的表，而后传入id，最后进行修改
	err = db.Model(&cate).Where("id = ?", id).Update(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//删除分类
func DeleteCategory(id int) int {
	var cate Category
	//删除分类的方法
	//TODO:为什么不限查询出来赋值给user而是直接可以链式调用删除user
	err = db.Debug().Where("id = ?", id).Delete(&cate).Error
	if err != nil { //说明有错
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
