package interfaces

import "net/http"

type FilterFactory interface {
	Apply(config interface{}) GatewayFilter
	GetOrder() int
}

type ServerWebExchange struct {
	Request *http.Request
}

type ResponseFilter func(*http.Response)

type ResponseFilters []ResponseFilter

func (this ResponseFilters) Filter(response *http.Response) {
	for _, filter := range this {
		filter(response)
	}
}

type GatewayFilter func(exchange *ServerWebExchange) ResponseFilter

func BuildServerWebExchange(request *http.Request) *ServerWebExchange {
	return &ServerWebExchange{Request: request}
}
