package v1

import (
	"gin_blog/model"
	"gin_blog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

/**
 * @Author: yirufeng
 * @Date: 2021/1/13 3:59 下午
 * @Desc:
 **/

//添加分类
func AddCategory(c *gin.Context) {
	//1. 拿到用户
	var cate model.Category

	//绑定模型
	//如果不需要冒黄我们就不需要接收
	_ = c.ShouldBindJSON(&cate)

	//2. 判断用户名是否存在
	code = model.CheckCategory(cate.Name)
	if code == errmsg.SUCCESS { //如果不存在就加入
		model.CreateCategory(&cate) // 写入数据库
	}
	//3.如果已经被使用了
	if code == errmsg.ERROR_CATEGORY_USED {
		code = errmsg.ERROR_CATEGORY_USED
	}

	//第一个200是网络传输的200，与我们业务操作状态码无关
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}


//TODO: 查询单个分类下的文章


//查询分类列表
func GetCategories(c *gin.Context) {
	//接收前端传过来的两个query
	pageSize, _ := strconv.Atoi(c.Query("pageSize")) //返回的是一个string
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))   //返回的是一个string

	//gorm官网：gorm.io/docs/query.html中有说到如果传入-1将会取消limit限制
	if pageSize == 0 {
		pageSize = -1
	}

	if pageNum == 0 {
		pageNum = -1
	}

	data := model.GetCategories(pageSize, pageNum)
	code = errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

//编辑分类
func EditCategory(c *gin.Context) {

	//获取用户名是否被占用
	var data model.Category
	c.ShouldBindJSON(&data)
	id, _ := strconv.Atoi(c.Param("id"))


	//TODO:这里也可以使用钩子函数改进，也就是在save之前check user
	code = model.CheckCategory(data.Name)
	if code == errmsg.SUCCESS {
		model.EditCategory(id, &data)
	}
	if code == errmsg.ERROR_CATEGORY_USED {
		c.Abort()
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

//删除分类
func DeleteCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteCategory(id)

	//可以直接返回,因为只需要一个成功或失败
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}


