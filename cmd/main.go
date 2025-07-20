package main

import (
	"fmt"
	"net"
	"os"

	"github.com/maruki00/zenithgo/internal/http/request"
	"github.com/maruki00/zenithgo/internal/server"
)

var _ = net.Listen
var _ = os.Exit

func main() {
	server := server.New()

	server.POST("/post", func(r *request.Request, a *any) {
		fmt.Println("hello world")
	})
	server.Run(":1234")

}
