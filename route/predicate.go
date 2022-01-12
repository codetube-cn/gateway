package route

import (
	"codetube.cn/gateway/predicate"
)

// Predicate 路由断言
type Predicate struct {
	Code      string
	Predicate predicate.Interface
}

// NewPredicate 创建路由断言
func NewPredicate() *Predicate {
	return &Predicate{}
}

// Predicates 路由断言列表
type Predicates []*Predicate
