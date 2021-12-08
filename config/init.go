package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

func InitConfig() *GatewayConfig {
	configFile := "gateway.yaml"
	configFileContent, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatal(err)
	}
	gatewayConfig := &GatewayConfig{}
	err = yaml.Unmarshal(configFileContent, gatewayConfig)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(gatewayConfig.Routes[0])
	return gatewayConfig
}
