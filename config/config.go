// 存放配置信息
package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"project-demo/utils/logs"
)

const (
	TimeFormat = "2006-01-02 15:04:05"
)

func initLogConfig(logLevel string) {
	if logLevel == "debug" {
		logrus.SetLevel(logrus.DebugLevel)
	} else if logLevel == "info" {
		logrus.SetLevel(logrus.InfoLevel)
	}
	//文件名和行号
	logrus.SetReportCaller(true)
	//日志格式json
	logrus.SetFormatter(&logrus.JSONFormatter{TimestampFormat: TimeFormat})
}

func init() {
	// 初始化配置
	logs.Debug(nil, "开始加载程序配置")
	viper.SetDefault("LOG_LEVEL", "info")
	viper.AutomaticEnv()
	logLevel := viper.GetString("LOG_LEVEL") //获取程序配置
	initLogConfig(logLevel)
}
