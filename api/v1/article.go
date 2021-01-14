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

//添加文章
func AddArticle(c *gin.Context) {
	//1. 拿到文章
	var article model.Article

	//绑定模型
	//如果不需要冒黄我们就不需要接收
	_ = c.ShouldBindJSON(&article)

	code = model.CreateArticle(&article)

	//第一个200是网络传输的200，与我们业务操作状态码无关
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    article,
		"message": errmsg.GetErrMsg(code),
	})
}

//TODO: 查询单个分类下的文章
func GetCategoryArticles(c *gin.Context) {
	//接收前端传过来的两个query
	pageSize, _ := strconv.Atoi(c.Query("pageSize")) //返回的是一个string
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))   //返回的是一个string
	id, _ := strconv.Atoi(c.Query("id"))             //返回的是一个string
	if pageSize == 0 {
		pageSize = -1
	}

	if pageNum == 0 {
		pageNum = -1
	}

	data, code := model.GetCategoryArticles(id, pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

//查询单个文章：当我们前端点击标题的时候，传递标题id给后端，后端查出来然后返回给前端进行渲染，
func GetArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := model.GetArticleInfo(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

//查询文章列表
func GetArticles(c *gin.Context) {
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

	var data []model.Article
	data, code = model.GetArticles(pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

//编辑文章
func EditArticle(c *gin.Context) {

	//获取用户名是否被占用
	var data model.Article
	id, _ := strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&data)

	code = model.EditArticle(id, &data)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

//删除文章
func DeleteArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteArticle(id)

	//可以直接返回,因为只需要一个成功或失败
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
