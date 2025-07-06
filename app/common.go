package main

import (
	"net"
	"os"
)

var _ = net.Listen
var _ = os.Exit

var Statues = map[int]string{
	//informational: 100's

	//Success: 200's
	200: "OK",
	201: "Created",

	//Redirect: 300's

	//Client Error: 400's
	404: "Not Found",

	//Server Errors: 500's
}

const CRLF = "\r\n"

var NOT_FOUND = []byte("HTTP/1.1 404 Not Found\r\n\r\n")

var AbsPath = "/tmp/"
