package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"runtime/pprof"

	"github.com/maruki00/azago/internal/azago"
)

var _ = net.Listen
var _ = os.Exit

func main() {

	f, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close() 
	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal(err)
	}
	defer pprof.StopCPUProfile() 
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
