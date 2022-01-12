package config

// GatewayConfig 网关配置
var GatewayConfig *Config

func init() {
	GatewayConfig = InitConfig()
}
