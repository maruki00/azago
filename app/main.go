package main

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"os"
	"path"
	"regexp"
	"strings"
)

var _ = net.Listen
var _ = os.Exit

var Statues = map[int]string{
	200: "OK",
	404: "Not Found",
}

const CRLF = "\r\n"

var NOT_FOUND = []byte("HTTP/1.1 404 Not Found\r\n\r\n")

type Response struct {
	r    *Request
	conn net.Conn
}

func NewResponse(r *Request, conn net.Conn) *Response {
	return &Response{
		r:    r,
		conn: conn,
	}
}

func (resp *Response) Write(status int, body []byte) {
	lenght := len(body)
	responseBody := strings.Builder{}
	responseBody.WriteString(fmt.Sprintf("HTTP/1.1 %d %s\r\n", status, Statues[status]))
	for header, value := range resp.r.Headers {
		responseBody.WriteString(fmt.Sprintf("%s: %s\r\n", header, value))
	}
	responseBody.WriteString("Content-Type: text/plain\r\n")
	responseBody.WriteString(fmt.Sprintf("Content-Length: %d\r\n", lenght))
	responseBody.WriteString("\r\n")
	responseBody.Write(body)
	resp.conn.Write([]byte(responseBody.String()))
}

func (resp *Response) SetHeaders(headers map[string]string) {
	// maps.Copy(headers, resp.r.Headers)
	for header, value := range headers {
		resp.r.Headers[header] = value
	}

}

type Request struct {
	Method      string
	UserAgent   string
	EndPoint    string
	HTTPVersion string
	Headers     map[string]string
}

func NewRequest() *Request {
	return &Request{
		Headers: make(map[string]string),
	}
}

func (req *Request) isValidEndPoint(rgx string) bool {
	regx, err := regexp.Compile(fmt.Sprintf("^%s$", rgx))
	if err != nil {
		return false
	}
	return regx.Match([]byte(req.EndPoint))
}
func RequestParser(conn net.Conn) (*Request, error) {
	request := NewRequest()
	requestBuff := make([]byte, 4096)
	_, err := conn.Read(requestBuff)
	if err != nil {
		panic(err)
	}
	requestInfo := bytes.Split(requestBuff, []byte(CRLF))
	if len(requestInfo) == 0 {
		return nil, errors.New("request is empty")
	}
	requestLine := bytes.Split(requestInfo[0], []byte(" "))
	if len(requestLine) == 0 {
		return nil, errors.New("invalide request line")
	}
	rLineLenght := len(requestLine)

	if rLineLenght >= 1 {
		request.Method = string(requestLine[0])
	}
	if rLineLenght >= 2 {
		request.EndPoint = string(requestLine[1])
	}
	if rLineLenght >= 3 {
		request.HTTPVersion = "HTTP/1.1" //string(requestLine[2])
	}
	if len(requestInfo) > 1 {
		for _, h := range requestInfo[1:] {
			header := bytes.Split(h, []byte(":"))
			if len(header) == 0 {
				continue
			}
			if len(bytes.Trim(header[0], " ")) == 0 {
				continue
			}
			if len(header) < 2 {
				request.Headers[string(bytes.Trim(header[0], " "))] = ""
				continue
			}
			request.Headers[string(bytes.Trim(header[0], " "))] = string(bytes.Trim(header[1], " "))
		}
	}
	return request, nil
}

func main() {
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

	if request.isValidEndPoint("/files/(.+)") {
		filesEndPoint(request, response)
		return
	}

	response.Write(404, []byte(Statues[404]))
	conn.Close()
}
func index(_ *Request, respone *Response) {
	respone.Write(200, []byte(Statues[200]))
}

func userAgentReaderEndPoint(request *Request, respone *Response) {
	if userAgent, ok := request.Headers["User-Agent"]; ok {
		respone.Write(200, []byte(userAgent))
		return
	}
	respone.Write(400, []byte(Statues[400]))
}
func echoEndPoint(request *Request, respone *Response) {
	query := strings.Replace(request.EndPoint, "/echo/", "", 6)
	respone.Write(200, []byte(query))
}
func filesEndPoint(request *Request, respone *Response) {
	filePath := strings.Replace(request.EndPoint, "/files/", "", 6)

	fullPath := path.Join("/tmp", filePath)
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		respone.Write(404, []byte(Statues[404]))
		return
	}
	contet, err := os.ReadFile(fullPath)
	if err != nil {
		respone.Write(404, []byte(Statues[404]))
		return
	}
	respone.SetHeaders(map[string]string{
		"Content-Type": "application/octet-stream",
	})
	fmt.Println("hellow orld")
	respone.Write(200, contet)
}
