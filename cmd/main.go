package main

import (
	"net"
	"os"

	"github.com/maruki00/zenithgo/internal/server"
)

var _ = net.Listen
var _ = os.Exit

func main() {
	server := server.New()


	server.Run(":1234")
}
