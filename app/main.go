package main

import (
	"fmt"
	"net"
	"os"
	"regexp"
	"strings"
)

// Ensures gofmt doesn't remove the "net" and "os" imports above (feel free to remove this!)
var _ = net.Listen
var _ = os.Exit

const CRLF = "\r\n"

func main() {
	listner, err := net.Listen("tcp", ":4221")
	if err != nil {
		panic(err)
	}

	conn, err := listner.Accept()
	if err != nil {
		panic(err)
	}

	var buff = make([]byte, 4048)

	_, err = conn.Read(buff)
	if err != nil {
		panic(err)
	}

	resp := []byte("HTTP/1.1 404 Not Found\r\n\r\n")
	request := string(buff)
	lines := strings.Split(request, CRLF)
	if len(lines) == 0 {
		conn.Write(resp)
		return
	}

	line1 := strings.Split(lines[0], " ")
	if len(line1) == 0 {
		conn.Write(resp)
		return
	}

	if line1[0] == "/" {
		conn.Write([]byte("HTTP/1.1 200 OK\r\n"))
		return
	}

	reg, err := regexp.Compile("/echo/(.)+")
	if err != nil {
		conn.Write(resp)
		return
	}

	if !reg.Match([]byte(line1[1])) {
		conn.Write(resp)
		return

	}

	path := strings.Replace(line1[1], "/echo/", "", 6)

	if len(path) == 0 { //&& line1[1] != "/" {
		conn.Write(resp)
		return
	}

	conn.Write([]byte(fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s", len(path), path)))
	conn.Close()
}
