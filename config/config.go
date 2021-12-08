package config

import "codetube.cn/gateway/route"

type Listen struct {
	Host string
	Port int
}

type GatewayConfig struct {
	Listen Listen
	Routes route.Routes
}
