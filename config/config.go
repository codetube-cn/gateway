package config

import "cn.codetube.gateway/route"

type Listen struct {
	Host string
	Port int
}

type GatewayConfig struct {
	Listen Listen
	Routes route.Routes
}
