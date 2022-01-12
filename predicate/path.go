package predicate

import (
	"encoding/json"
	"net/http"
	"path/filepath"
)

// PathPredicate path 断言
type PathPredicate struct {
	Contract
	Value        string //值
	DefaultValue string //默认值
}

// NewPathPredicate 创建 path 断言
func NewPathPredicate() Interface {
	return &PathPredicate{Value: "", DefaultValue: ""}
}

//LoadValue 载入断言值，参数一般为 json
func (p *PathPredicate) LoadValue(v string) error {
	return json.Unmarshal([]byte(v), &p.Value)
}

//Match 匹配断言
func (p *PathPredicate) Match(request *http.Request) bool {
	matched, err := filepath.Match(p.Value, request.URL.Path)
	if err != nil || !matched {
		return false
	}
	return true
}
