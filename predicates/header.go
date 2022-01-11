package predicates

import (
	"encoding/json"
	"net/http"
	"regexp"
)

type HeaderPredicate struct {
	PredicateContract
	Value        []HeaderPredicateValue //值
	DefaultValue []HeaderPredicateValue //默认值
}

func NewHeaderPredicate() PredicateInterface {
	return &HeaderPredicate{}
}

type HeaderPredicateValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (p *HeaderPredicate) LoadValue(v string) error {
	return json.Unmarshal([]byte(v), &p.Value)
}

func (p *HeaderPredicate) Match(request *http.Request) bool {
	//逐一进行匹配
	//header 中无 key，匹配失败，有 key 而不能匹配值（支持正则），匹配失败
	for _, h := range p.Value {
		if value, ok := request.Header[h.Key]; !ok {
			return false
		} else {
			//如果为正则表达式，需要匹配
			reg, err := regexp.Compile(h.Value)
			if err != nil {
				return false
			}
			//非正则表达式，则按字符串完全匹配
			if !reg.MatchString(value[0]) {
				return false
			}
		}
	}

	return true
}
