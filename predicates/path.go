package predicates

import (
	"encoding/json"
	"net/http"
	"path/filepath"
)

type PathPredicate struct {
	PredicateContract
	Value        string //值
	DefaultValue string //默认值
}

func NewPathPredicate() PredicateInterface {
	return &PathPredicate{Value: "", DefaultValue: ""}
}

func (p *PathPredicate) LoadValue(v string) error {
	return json.Unmarshal([]byte(v), &p.Value)
}

func (p *PathPredicate) Match(request *http.Request) bool {
	matched, err := filepath.Match(p.Value, request.URL.Path)
	if err != nil || !matched {
		return false
	}
	return true
}
