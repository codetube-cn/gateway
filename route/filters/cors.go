package filters

import (
	"codetube.cn/gateway/interfaces"
	"net/http"
)

func init() {
	RegisterFilter("CORS", NewCORSFilter())
}

type CORSFilter struct{}

func NewCORSFilter() *CORSFilter {
	return &CORSFilter{}
}

func (f *CORSFilter) Apply(config interface{}) interfaces.GatewayFilter {
	return func(exchange *interfaces.ServerWebExchange) interfaces.ResponseFilter {
		return func(response *http.Response) {
			response.Header.Add("Access-Control-Allow-Origin", "*")
			response.Header.Add("Access-Control-Allow-Headers", "Origin, Content-Type, Cookie,X-CSRF-TOKEN, Accept,Authorization")
			response.Header.Add("Access-Control-Expose-Headers", "Authorization,authenticated")
			response.Header.Add("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, OPTIONS, DELETE")
			response.Header.Add("Access-Control-Allow-Credentials", "true")
		}
	}
}

func (f *CORSFilter) GetOrder() int {
	return -1
}


