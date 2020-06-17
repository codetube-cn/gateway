package filters

import (
	"cn.codetube.gateway/interfaces"
	"net/http"
)

func init() {
	RegisterFilter("CORS", NewCORSFilter())
}

type CORSFilter struct{}

func (this *CORSFilter) Apply(config interface{}) interfaces.GatewayFilter {
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

func (this *CORSFilter) GetOrder() int {
	return -1
}

func NewCORSFilter() *CORSFilter {
	return &CORSFilter{}
}
