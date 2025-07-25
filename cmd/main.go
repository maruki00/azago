package main

import (
	"net"
	"os"

	"github.com/maruki00/azago/internal/azago"
	"github.com/maruki00/azago/internal/server"
)

var _ = net.Listen
var _ = os.Exit

func main() {
	server := server.New()

	server.POST("/ping/:id", func(r *azago.Context) {
		r.WriteJSON(200, "hello world")

	})
	server.Run(":1234")

}
