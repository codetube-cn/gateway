package filters

import (
	"codetube.cn/gateway/interfaces"
)

func init() {
	RegisterFilter("SetHeader", NewSetHeaderFilter())
}

type SetHeaderFilter struct{}

func (this *SetHeaderFilter) Apply(config interface{}) interfaces.GatewayFilter {
	return func(exchange *interfaces.ServerWebExchange) interfaces.ResponseFilter {
		p := NameConfig(config.(string))
		if headers := p.GetValue(); headers != nil {
			for _, header := range headers {
				exchange.Request.Header.Set(header.Name, header.Value)
			}
		}
		return nil
	}
}

func (this *SetHeaderFilter) GetOrder() int {
	return 1
}

func NewSetHeaderFilter() *SetHeaderFilter {
	return &SetHeaderFilter{}
}
