package filters

import (
	"cn.codetube.gateway/interfaces"
	"sync"
)

var FilterMap sync.Map

func RegisterFilter(name string, value interfaces.FilterFactory) {
	FilterMap.Store(name, value)
}
