package predicate

import (
	"encoding/json"
	"net/http"
	"strings"
)

// MethodPredicate method 断言
type MethodPredicate struct {
	PredicateContract
	Value        []string //值
	DefaultValue []string //默认值
}

// NewMethodPredicate 创建 method 断言
func NewMethodPredicate() PredicateInterface {
	return &MethodPredicate{}
}

//LoadValue 载入断言值，参数一般为 json
func (p *MethodPredicate) LoadValue(v string) error {
	return json.Unmarshal([]byte(v), &p.Value)
}

//Match 匹配断言
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
