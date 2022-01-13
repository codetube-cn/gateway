package predicate

import (
	"encoding/json"
	"net/http"
	"regexp"
)

// HeaderPredicate header 断言
type HeaderPredicate struct {
	Contract
	Value        []HeaderPredicateValue //值
	DefaultValue []HeaderPredicateValue //默认值
}

// NewHeaderPredicate 创建 header 断言
func NewHeaderPredicate() Interface {
	return &HeaderPredicate{}
}

// HeaderPredicateValue 断言值
type HeaderPredicateValue struct {
	Key           string `json:"key"`            // header key，要注意大小写
	Value         string `json:"value"`          //header value
	MatchType     string `json:"match_type"`     //匹配模式，完全匹配|must、包含|include
	CaseSensitive bool   `json:"case_sensitive"` //是否大小写敏感
}

//LoadValue 载入断言值，参数一般为 json
func (p *HeaderPredicate) LoadValue(v string) error {
	return json.Unmarshal([]byte(v), &p.Value)
}

//Match 匹配断言
func (p *HeaderPredicate) Match(request *http.Request) bool {
	//逐一进行匹配
	//header 中无 key，匹配失败，有 key 而不能匹配值（支持正则），匹配失败
	for _, h := range p.Value {
		if value, ok := request.Header[h.Key]; !ok {
			return false
		} else {
			pattern := h.Value
			if h.MatchType == "full" {
				pattern = "^" + pattern + "$"
			}
			if !h.CaseSensitive {
				pattern = "(?i)" + pattern
			}
			//如果为正则表达式，需要匹配
			reg, err := regexp.Compile(pattern)
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
