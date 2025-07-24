// Package server
package server

import (
	"fmt"
	"net"
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

func (_this *Server) HandleRequest(conn net.Conn) {
	start := time.Now()
	defer conn.Close()
	request := azago.NewRequest(conn)
	if request == nil {
		conn.Write([]byte(common.INTERNAL_ERROR))
		return
	}

	response := azago.NewResponse(request, conn)
	// routes := _this.GetRoutes()
	// for k, r := range routes {
	// 	fmt.Printf("route : %s, %+v, %+v\n", k, r, r.Params)
	// }
	route := _this.GetRoute(request.EndPoint)
	fmt.Println(route)

	response.Write(200, []byte("hello worlld"))
	logPkg.Log("spent", timePkg.Since(start), request.EndPoint)
}
