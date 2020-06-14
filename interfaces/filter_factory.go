package interfaces

import "net/http"

type FilterFactory interface {
	Apply() GatewayFilter
}

type ServerWebExchange struct {
	Request *http.Request
}

type GatewayFilter func(exchange *ServerWebExchange)

func BuildServerWebExchange(request *http.Request) *ServerWebExchange {
	return &ServerWebExchange{Request: request}
}
