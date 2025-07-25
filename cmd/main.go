package main

import (
	"fmt"
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
		fmt.Println("hello world")
	})
	server.Run(":1234")

}
