package config

// GatewayConfig 网关配置
var GatewayConfig *Config

func init() {
	//初始化配置
	GatewayConfig = InitConfig()
}
