package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) >= 3 {
		AbsPath = os.Args[2]
	}
	listner, err := net.Listen("tcp", ":4221")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listner.Accept()
		if err != nil {
			panic(err)
		}
		go handelRequest(conn)
	}
}

func handelRequest(conn net.Conn) {
	request, err := RequestParser(conn)
	if err != nil {
		conn.Write(NOT_FOUND)
		return
	}
	defer conn.Close()
	response := NewResponse(request, conn)
	if request.EndPoint == "/" {
		index(request, response)
		return
	}

	if request.EndPoint == "/user-agent" {
		userAgentReaderEndPoint(request, response)
		return
	}

	if request.isValidEndPoint("/echo/(.+)") {
		echoEndPoint(request, response)
		return
	}

	if request.isValidEndPoint("/files/(.+)") && request.Method == "GET" {
		fmt.Println("request get recieved")
		getFilesEndPoint(request, response)
		return
	}

	if request.isValidEndPoint("/files/(.+)") && request.Method == "POST" {
		fmt.Println("request post recieved")
		postFilesEndPoint(request, response)
		return
	}

	response.Write(404, []byte(Statues[404]))
	conn.Close()
}
