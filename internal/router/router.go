// this is just a simple router
// missing groupes and many other stuffs
// middlewares by groups ...
package router

import "github.com/maruki00/zenithgo/internal/http/request"

type HttpHandler func(*request.Request, *any)
type HttpMiddleware func()
type Route struct {
	Handler HttpHandler
	Method string
	Params map[string]any
	Middlewares []HttpMiddleware
}
type Routes map[string]*Route
type Router struct {
	Middlewares []HttpMiddleware
	routes Routes
}
func NewRouter() *Router {
	return &Router{
		Middlewares: make([]HttpMiddleware, 0),
		routes: make(Routes),
	}
}
func (_this *Router) GetRoutes() Routes {
	return _this.routes
}
func (_this*Router) Add(method string, pattern string, handler HttpHandler, middlewares ...HttpMiddleware) {
	middle_wares := make([]HttpMiddleware, len(_this.Middlewares)+len(middlewares))
	middle_wares = append(middle_wares, _this.Middlewares...)
	middle_wares = append(middle_wares, middlewares...)
	_this.routes[pattern] = &Route{
		Handler: handler,
		Method: method,
		Middlewares: middle_wares,
	}
}
func (_this *Router) POST(pattern string, handler HttpHandler, middlewares ...HttpMiddleware) {
	_this.Add("POST", pattern , handler, middlewares...)
}
func (_this *Router) GET(pattern string, handler HttpHandler, middlewares ...HttpMiddleware) {
	_this.Add("GET", pattern , handler, middlewares...)
}

