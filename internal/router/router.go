// Package router
// this is just a simple router
// missing groupes and many other stuffs
// middlewares by groups ...
package router

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/maruki00/azago/internal/azago"
)

type HTTPHandler func(*azago.Request, *any)
type HTTPMiddleware func()
type Route struct {
	Handler     HTTPHandler
	Method      string
	//uint8 because usualy params not much
	ParamNames  map[string]uint8
	Params      map[string]string
	Middlewares []HTTPMiddleware
}
type Routes map[string]*Route
type Router struct {
	Middlewares []HTTPMiddleware
	routes      Routes
}

func NewRouter() *Router {
	return &Router{
		Middlewares: make([]HTTPMiddleware, 0),
		routes:      make(Routes),
	}
}

func (_this *Router) GetEndPoint(route string) {
	var endpoint strings.Builder
	for pattern := range _this.routes {
		endpoint.Reset()
		endpoint.WriteString(pattern)

	}
}

func (_this *Router) GetRoutes() Routes {
	return _this.routes
}

func (_this *Router) Add(method string, pattern string, handler HTTPHandler, middlewares ...HTTPMiddleware) {
	_midllewares := make([]HTTPMiddleware, len(_this.Middlewares)+len(middlewares))
	_midllewares = append(_midllewares, _this.Middlewares...)
	_midllewares = append(_midllewares, middlewares...)
	route := &Route{
		Handler:     handler,
		Method:      method,
		Middlewares: _midllewares,
		ParamNames:  make(map[string]uint8),
		Params:      make(map[string]string, 0),
	}
	parts := strings.Split(strings.Trim(pattern, "/"), "/")
	fmt.Println("parts : ", parts)
	var prefix []rune
	for i, part := range parts {
		if part == "" {
			continue
		}
		prefix = []rune(part)
		if len(prefix) < 2 || prefix[0] != ':' {
			continue
		}
		route.ParamNames[part[1:]] = uint8(i)
		route.Params[part[1:]] = ""
		parts[i] = "(.+)"
	}
	_this.routes[strings.Join(parts, "/")] = route
}

func (_this *Router) POST(pattern string, handler HTTPHandler, middlewares ...HTTPMiddleware) {
	_this.Add("POST", pattern, handler, middlewares...)
}

func (_this *Router) GET(pattern string, handler HTTPHandler, middlewares ...HTTPMiddleware) {
	_this.Add("GET", pattern, handler, middlewares...)
}

func (_this *Router) GetRoute(pattern string) *Route {
	var rgx *regexp.Regexp
	var err error
	for k, v := range _this.routes {
		rgx, err = regexp.Compile(k)
		if err != nil {
			continue
		}
		if rgx.Match([]byte(pattern)) {
			return v
		}
	}
	return nil
}
