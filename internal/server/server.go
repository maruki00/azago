// Package server
package server

import (
	"context"
	"net"
	"strings"
	"time"

	"github.com/maruki00/azago/internal/azago"
	"github.com/maruki00/azago/internal/common"
	"github.com/maruki00/azago/internal/router"
	logPkg "github.com/maruki00/azago/pkg/log"
	timePkg "github.com/maruki00/azago/pkg/time"
)

type Server struct {
	*router.Router
}

func New() *Server {
	return &Server{
		Router: router.NewRouter(),
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

var ctx *azago.Context

func (_this *Server) HandleRequest(conn net.Conn) {
	start := time.Now()
	defer conn.Close()
	request := azago.NewRequest(conn)
	if request == nil {
		conn.Write([]byte(common.INTERNAL_ERROR))
		return
	}

	response := azago.NewResponse(request, conn)
	route := _this.GetRoute(request.EndPoint)
	parts := strings.Split(strings.Trim(request.EndPoint, "/"), "/")

	for k, v := range route.ParamNames {
		request.Params[k] = parts[v]
	}
	ctx = &azago.Context{
		Context:  context.Background(),
		Request:  request,
		Response: response,
	}
	route.Handler(ctx)
	logPkg.Log(request.Method, timePkg.Since(start), request.EndPoint)
}
