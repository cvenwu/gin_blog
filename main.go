package main

import (
	"gin_blog/model"
	"gin_blog/routes"
)

/**
 * @Author: yirufeng
 * @Date: 2021/1/11 4:14 下午
 * @Desc:
 **/

func main() {
	model.InitDB()
	routes.InitRouter()
}