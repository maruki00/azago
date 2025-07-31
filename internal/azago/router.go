package azago

type HTTPHandler func(ctx *Context)
type HTTPMiddleware func(ctx *Context) error
type Route struct {
	Part       string
	Handler    map[string]HTTPHandler
	Childs     []*Route
	IsParam    bool
	IsWildCard bool
}
func NewRoute(Part string, isPrm,isWCard bool)*Route{
	return &Route{
		Handler: make(map[string]HTTPHandler),
		Childs: make([]*Route,0,1),
		Part: Part,
		IsParam: isPrm,
		IsWildCard: isWCard,
	}
}
type Router struct {
	Root *Route
}
func NewRouter() *Router {
	return nil
}

func (_this *Router) GetEndPoint(route string) {
}

func (_this *Router) GetRoutes() any {
	return nil
}

func (_this *Router) Add(method string, pattern string, handler HTTPHandler, middlewares ...HTTPMiddleware) {
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





// type Route struct {
// 	Handler     HTTPHandler
// 	Method      string
// 	ParamNames  map[string]uint8
// 	Middlewares []HTTPMiddleware
// }
// type Routes map[string]*Route
// type Router struct {
// 	Middlewares []HTTPMiddleware
// 	routes      Routes
// }
//
// func NewRouter() *Router {
// 	return &Router{
// 		Middlewares: make([]HTTPMiddleware, 0),
// 		routes:      make(Routes),
// 	}
// }
//
// func (_this *Router) GetEndPoint(route string) {
// 	var endpoint strings.Builder
// 	for pattern := range _this.routes {
// 		endpoint.Reset()
// 		endpoint.WriteString(pattern)
//
// 	}
// }
//
// func (_this *Router) GetRoutes() Routes {
// 	return _this.routes
// }
//
// func (_this *Router) Add(method string, pattern string, handler HTTPHandler, middlewares ...HTTPMiddleware) {
// 	_midllewares := make([]HTTPMiddleware, len(_this.Middlewares)+len(middlewares))
// 	_midllewares = append(_midllewares, _this.Middlewares...)
// 	_midllewares = append(_midllewares, middlewares...)
// 	route := &Route{
// 		Handler:     handler,
// 		Method:      method,
// 		Middlewares: _midllewares,
// 		ParamNames:  make(map[string]uint8),
// 	}
// 	parts := strings.Split(strings.Trim(pattern, "/"), "/")
// 	var prefix []rune
// 	for i, part := range parts {
// 		if part == "" {
// 			continue
// 		}
// 		prefix = []rune(part)
// 		if len(prefix) < 2 || prefix[0] != ':' {
// 			continue
// 		}
// 		route.ParamNames[part[1:]] = uint8(i)
// 		parts[i] = "(.+)"
// 	}
// 	_this.routes[strings.Join(parts, "/")] = route
// }
//
// func (_this *Router) POST(pattern string, handler HTTPHandler, middlewares ...HTTPMiddleware) {
// 	_this.Add("POST", pattern, handler, middlewares...)
// }
//
// func (_this *Router) GET(pattern string, handler HTTPHandler, middlewares ...HTTPMiddleware) {
// 	_this.Add("GET", pattern, handler, middlewares...)
// }
//
// func (_this *Router) GetRoute(pattern string) *Route {
// 	var rgx *regexp.Regexp
// 	var err error
// 	for k, v := range _this.routes {
// 		rgx, err = regexp.Compile(k)
// 		if err != nil {
// 			continue
// 		}
// 		if rgx.Match([]byte(pattern)) {
// 			return v
// 		}
// 	}
// 	return nil
// }
