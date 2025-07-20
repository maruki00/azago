package server

import (
	"net"
	"time"

	"github.com/maruki00/zenithgo/internal/common"
	"github.com/maruki00/zenithgo/internal/http/request"
	"github.com/maruki00/zenithgo/internal/http/response"
	"github.com/maruki00/zenithgo/internal/router"
	logPkg "github.com/maruki00/zenithgo/pkg/log"
	timePkg "github.com/maruki00/zenithgo/pkg/time"
)

type Server struct {
	router.Router
}

func New() *Server {
	return &Server{}
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
		go _this.NewRequest(conn)
	}
}

func (_this *Server) NewRequest(conn net.Conn) {
	start := time.Now()
	defer conn.Close()
	request := request.NewRequest(conn)
	if request == nil {
		conn.Write([]byte(common.INTERNAL_ERROR))
		return
	}

	response := response.NewResponse(request, conn)

	response.Write(200, []byte("hello worlld"))
	logPkg.Log("spent", timePkg.Since(start), request.EndPoint)
}
