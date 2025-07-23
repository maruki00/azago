package main

import (
	"fmt"
	"net"
	"os"

)

var _ = net.Listen
var _ = os.Exit

func main() {
	server := server.New()

	server.POST("/ping/:id", func(r *request.Request, a *any) {
		fmt.Println("hello world")
	})
	server.Run(":1234")

}
