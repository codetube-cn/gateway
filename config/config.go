package config

import "codetube.cn/gateway/route"

type Listen struct {
	Host string
	Port int
}

type GatewayConfig struct {
	Gateway string
	Listen Listen
	Routes route.Routes
}
