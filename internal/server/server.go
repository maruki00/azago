package server

import (
	"log/slog"
	"net"
	"time"

	"github.com/maruki00/zenithgo/internal/common"
	"github.com/maruki00/zenithgo/internal/http/request"
	"github.com/maruki00/zenithgo/internal/http/response"
	"github.com/maruki00/zenithgo/internal/router"
	timePkg "github.com/maruki00/zenithgo/pkg/time"
)

type Server struct {
	router.Router
}

func New() *Server {
	return &Server{}
}

func (_this *Server) Run(addr string) {
	slog.Info("Server start at", "addr ", addr)
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
	start := time.Now()
	defer conn.Close()
	go func() {
		request := request.NewRequest(conn)
		if request == nil {
			conn.Write([]byte(common.INTERNAL_ERROR))
			return
		}
		_ = response.NewResponse(request, conn)
	}()
	slog.Info("request", "spend", timePkg.Since(start), "from", conn.LocalAddr().String())
}
