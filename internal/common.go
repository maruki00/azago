package main



var responses = map[string][]byte {
	"OK" : []byte("HTTP/1.1 200 OK\r\n\r\n"),
	"NOT_FOUND" : []byte("HTTP/1.1 404 Not Found\r\n\r\n"),
}



var AbsPath = "/tmp/"
const CRLF = "\r\n"
