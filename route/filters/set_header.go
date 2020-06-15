package filters

import "cn.codetube.gateway/interfaces"

func init() {
	RegisterFilter("SetHeader", NewSetHeaderFilter())
}

type SetHeaderFilter struct{}

func (this *SetHeaderFilter) Apply(config interface{}) interfaces.GatewayFilter {
	return func(exchange *interfaces.ServerWebExchange) {
		p := NameConfig(config.(string))
		if headers := p.GetValue(); headers != nil {
			for _, header := range headers {
				exchange.Request.Header.Set(header.Name, header.Value)
			}
		}
	}
}

func NewSetHeaderFilter() *SetHeaderFilter {
	return &SetHeaderFilter{}
}
