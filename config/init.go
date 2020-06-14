package config

import (
	"github.com/micro/go-micro/v2/config"
	"log"
)

func InitConfig() GatewayConfig {
	configFile := "gateway.yaml"
	err := config.LoadFile(configFile)
	if err != nil {
		log.Fatal(err)
	}
	gatewayConfig := GatewayConfig{}
	err = config.Scan(&gatewayConfig)
	if err != nil {
		log.Fatal(err)
	}
	return gatewayConfig
}
