//项目的总路口

package main

import (
	"github.com/gin-gonic/gin"
	_ "project-demo/config"
)

func main() {
	//1. 初始化配置
	//2.配置gin
	r := gin.Default()
	r.Run()
}
