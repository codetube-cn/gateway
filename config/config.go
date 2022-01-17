package config

import (
	"codetube.cn/core/config"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

// Listen 网关监听配置
type Listen struct {
	Host string //监听主机
	Port int    //监听端口
}

// Config 网关配置
type Config struct {
	Gateway              string              `yaml:"gateway"`                //网关标识
	WatchIntervalSeconds int                 `yaml:"watch-interval-seconds"` //监听配置变化时间间隔（秒）
	Listen               *Listen             `yaml:"listen"`                 //网关监听配置
	Mysql                *config.MysqlConfig `yaml:"mysql"`                  //数据库连接配置
}

// NewConfig 创建网关配置
func NewConfig() *Config {
	return &Config{}
}

// InitConfig 初始化配置
func InitConfig() *Config {
	configFile := "config.yaml"
	configFileContent, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatal(err)
	}
	gatewayConfig := NewConfig()
	err = yaml.Unmarshal(configFileContent, gatewayConfig)
	if err != nil {
		log.Fatal(err)
	}
	if gatewayConfig.Gateway == "" {
		log.Fatal("config err: invalid gateway")
	}
	return gatewayConfig
}
