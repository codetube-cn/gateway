package filters

import (
	"codetube.cn/gateway/interfaces"
	"strings"
	"sync"
)

var FilterMap sync.Map

func RegisterFilter(name string, value interfaces.FilterFactory) {
	FilterMap.Store(name, value)
}

type ValueConfig string

func (this ValueConfig) GetValue() []string {
	return strings.Split(string(this), ",")
}

type NameConfig string
type NameConfigObj struct {
	Name  string
	Value string
}

func (this NameConfig) GetValue() []*NameConfigObj {
	slist := strings.Split(string(this), ",")
	if len(slist) < 2 || len(slist)%2 != 0 {
		return nil
	}
	result := make([]*NameConfigObj, 0)
	for i := 0; i < len(slist); i = i + 2 {
		result = append(result, &NameConfigObj{
			Name:  slist[i],
			Value: slist[i+1],
		})
	}
	return result
}
