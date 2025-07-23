// Package server
package server


import (
	"fmt"
	"net"
	"time"

	"github.com/maruki00/azago/internal/common"
	"github.com/maruki00/azago/internal/http/request"
	"github.com/maruki00/azago/internal/http/response"
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
	request := request.NewRequest(conn)
	if request == nil {
		conn.Write([]byte(common.INTERNAL_ERROR))
		return
	}

	response := response.NewResponse(request, conn)
	routes := _this.GetRoutes()
	// for k, r := range routes {
	// 	fmt.Printf("route : %s, %+v", k, r)
	// }
	fmt.Println(routes)

	response.Write(200, []byte("hello worlld"))
	logPkg.Log("spent", timePkg.Since(start), request.EndPoint)
}
