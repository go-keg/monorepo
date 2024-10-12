package middleware

import (
	"net/http"
	"regexp"
	"strings"
)

type Routes []string

func NewRoutes(items []string) *Routes {
	var routes = make(Routes, len(items))
	for i, route := range items {
		routes[i] = strings.ReplaceAll(route, " ", "")
	}
	return &routes
}

func (r Routes) InRoutes(req *http.Request) bool {
	for _, route := range r {
		if route != "" && matchRoute(route, req.Method+req.URL.Path) {
			return true
		}
	}
	return false
}

func (r Routes) Has() bool {
	return len(r) > 0
}

var exp = regexp.MustCompile(`/\{[\w-]+}`)

func matchRoute(route, path string) bool {
	route = exp.ReplaceAllLiteralString(route, "/[\\w-]+") + "$"
	if ok, _ := regexp.Match(route, []byte(path)); ok {
		return true
	}
	return false
}
