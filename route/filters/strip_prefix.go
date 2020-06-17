package filters

import (
	"cn.codetube.gateway/interfaces"
	"log"
	"strconv"
	"strings"
)

func init() {
	RegisterFilter("StripPrefix", NewStripPrefixFilter())
}

type StripPrefixFilter struct{}

func (this *StripPrefixFilter) Apply(config interface{}) interfaces.GatewayFilter {
	return func(exchange *interfaces.ServerWebExchange) interfaces.ResponseFilter {
		path := exchange.Request.URL.Path
		defaultIndex := 1
		config := ValueConfig(config.(string))
		i, err := strconv.Atoi(config.GetValue()[0])
		if err != nil {
			log.Println(err)
		} else {
			defaultIndex = i
		}
		pathList := strings.Split(path, "/")
		exchange.Request.URL.Path = strings.Join(pathList[defaultIndex+1:], "/")
		return nil
	}
}

func (this *StripPrefixFilter) GetOrder() int {
	return 2
}

func NewStripPrefixFilter() *StripPrefixFilter {
	return &StripPrefixFilter{}
}
