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
		var obj interface{}
		if err := r.BindJSON(obj); err != nil {
			r.WriteJSON(400, map[string]string{
				"status": "200",
				"message": err.Error(),
			})
			return
		}
		id,_ := r.Params["id"]
		fmt.Println("dev : ", id, obj)
		r.WriteJSON(200, "hello world")

	})
	server.Run(":1234")

}
