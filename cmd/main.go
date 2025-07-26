package main

import (
	"fmt"
	"net"
	"os"

	"github.com/maruki00/azago/internal/azago"
	"github.com/pkg/profile"
	"go.uber.org/automaxprocs/maxprocs"
)

var _ = net.Listen
var _ = os.Exit

func main() {
	_,err := maxprocs.Set()
	if err != nil {
		panic(err)
	}
	defer profile.Start(profile.TraceProfile, profile.ProfilePath(".")).Stop()

	server := azago.New()

	server.POST("/ping/:id", func(r *azago.Context) {
		var obj = make(map[string]any)
		if err := r.BindJSON(obj); err != nil {
			r.WriteJSON(400, map[string]string{
				"status":  "200",
				"message": err.Error(),
			})
			return
		}
		id, _ := r.Params["id"]
		fmt.Println("dev : ", id, obj)
		r.WriteJSON(200, "hello world")

	})
	server.Run(":1234")

}
