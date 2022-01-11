package predicates

import (
	"encoding/json"
	"net/http"
	"strings"
)

type MethodPredicate struct {
	PredicateContract
	Value        []string //值
	DefaultValue []string //默认值
}

func NewMethodPredicate() PredicateInterface {
	return &MethodPredicate{}
}

func (p *MethodPredicate) LoadValue(v string) error {
	return json.Unmarshal([]byte(v), &p.Value)
}

func (p *MethodPredicate) Match(request *http.Request) bool {
	//无值列表，为任意匹配
	if len(p.Value) == 0 {
		return true
	}

	//指定了 Method 列表，则需要匹配其中之一
	for _, method := range p.Value {
		if strings.ToUpper(method) == strings.ToUpper(request.Method) {
			return true
		}
	}

	//都未匹配上，匹配失败
	return false
}
