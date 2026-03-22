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

var (
	Port       string
	JwtSignKey string
	JwtExpTime int64 //token有效期
	Username   string
	Password   string
)

type ReturnData struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

// 构造函数，给结构体赋默认值
func NewReturnData() ReturnData {
	returnData := ReturnData{}
	returnData.Status = 200
	data := map[string]interface{}{}
	returnData.Data = data
	return returnData
}
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
	viper.SetDefault("LOG_LEVEL", "debug")
	//获取端口号配置
	viper.SetDefault("PORT", ":8081")
	Port = viper.GetString("PORT")
	//获取jwt加密的secret
	viper.SetDefault("JWT_SIGN_KEY", "123456")
	JwtSignKey = viper.GetString("JWT_SIGN_KEY")
	//获取jwt有效期配置
	viper.SetDefault("JWT_EXPIRE_TIME", 24*60*60) //24小时
	JwtExpTime = viper.GetInt64("JWT_EXPIRE_TIME")

	//配置用户名密码的默认值，生产环境是加密状态
	viper.SetDefault("USERNAME", "admin")
	viper.SetDefault("PASSWORD", "123456")

	//获取用户名和密码配置
	Username = viper.GetString("USERNAME")
	Password = viper.GetString("PASSWORD")
	//加载环境变量
	viper.AutomaticEnv()
	logLevel := viper.GetString("LOG_LEVEL") //获取程序配置
	initLogConfig(logLevel)

}
