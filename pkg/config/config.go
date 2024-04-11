package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// NewConfig 创建一个新的Viper配置实例
// p string 配置文件的路径
func NewConfig(p string) *viper.Viper {
	// 从环境变量中获取配置文件的路径
	envConf := os.Getenv("APP_CONF")
	// 如果环境变量中没有获取到，则使用传入的参数
	if envConf == "" {
		envConf = p
	}
	// 打印加载的配置文件路径
	fmt.Println("load conf file:", envConf)
	// 根据传入的参数加载配置文件并返回
	return getConfig(envConf)
}

func getConfig(path string) *viper.Viper {
	conf := viper.New()
	conf.SetConfigFile(path)
	err := conf.ReadInConfig()
	if err != nil {
		panic(err)
	}
	return conf
}
