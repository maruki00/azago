package server

import (
	"fmt"
	"net"

	"github.com/maruki00/zenithgo/internal/common"
	"github.com/maruki00/zenithgo/internal/http/request"
	"github.com/maruki00/zenithgo/internal/http/response"
	"github.com/maruki00/zenithgo/internal/router"
)

type Server struct {
	router.Router
}

func New() *Server {
	return &Server{}
}

func (_this *Server) Run(addr string) {
	listner, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listner.Accept()
		if err != nil {
			panic(err)
		}
		_this.NewRequest(conn)
	}
}

func (_this *Server) NewRequest(conn net.Conn) {
	defer conn.Close()
	request := request.NewRequest(conn)
	if request == nil {
		conn.Write([]byte(common.INTERNAL_ERROR))
		return
	}
	response := response.NewResponse(request, conn)
	

	response.Write(404, []byte(Statues[404]))
}

