package route

import (
	"cn.codetube.gateway/route/predicates"
)

type Predicates struct {
	Header predicates.HeaderPredicate
	Method predicates.MethodPredicate
	Host   string
	Path   predicates.PathPredicate
}

type Route struct {
	Id         string
	Url        string
	Predicates Predicates
	Filters    []interface{}
}

type Routes []*Route
