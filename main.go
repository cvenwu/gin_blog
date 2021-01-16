package main

import (
	"gin_blog/config"
	"github.com/gin-gonic/gin"
)

/**
 * @Author: yirufeng
 * @Date: 2021/1/16 10:25 上午
 * @Desc:
 **/

func main() {
	config.InitConfig()
	r := gin.Default()
	r.Run(":3000")
}
