package config


type Listen struct {
	Host string
	Port int
}

type GatewayConfig struct {
	Gateway string
	Listen Listen
}
