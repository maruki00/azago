package azago

import (
	"context"
	"net"
	"strings"
	"time"

	httpPkg "github.com/maruki00/azago/pkg/http"
	logPkg "github.com/maruki00/azago/pkg/log"
	timePkg "github.com/maruki00/azago/pkg/time"
)

type Server struct {
	*Router
}

func New() *Server {
	return &Server{
		Router: NewRouter(),
	}
}

func (_this *Server) Run(addr string) {
	logPkg.Log("Server start at", "addr ", addr)
	listner, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listner.Accept()
		if err != nil {
			panic(err)
		}
		go _this.HandleRequest(conn)
	}
}

var ctx *Context

func (_this *Server) HandleRequest(conn net.Conn) {
	start := time.Now()
	defer conn.Close()
	request := NewRequest(conn)
	if request == nil {
		conn.Write([]byte(httpPkg.INTERNAL_ERROR))
		return
	}

	response := NewResponse(request, conn)
	route := _this.GetRoute(request.EndPoint)
	if route.Method != request.Method {
		response.Write(httpPkg.StatusNotFound, []byte("Route Not Found"))
	}
	parts := strings.Split(strings.Trim(request.EndPoint, "/"), "/")

	for k, v := range route.ParamNames {
		request.Params[k] = parts[v]
	}

	ctx = &Context{
		Context:  context.Background(),
		Request:  request,
		Response: response,
	}

	var err error
	for _, mWare := range _this.Middlewares {
		err = mWare(ctx)
		if err != nil {
			break
		}
	}

	if err != nil {
		return
	}

	route.Handler(ctx)

	logPkg.Log(request.Method, timePkg.Since(start), request.EndPoint)
}
