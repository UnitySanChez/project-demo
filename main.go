//项目的总路口

package main

import (
	"github.com/gin-gonic/gin"
	"project-demo/config"

	"project-demo/utils/logs"
)

func main() {
	//1. 初始化配置
	//2.配置gin
	r := gin.Default()
	logs.Info(nil, "项目启动成功")
	r.Run(config.Port)

}
