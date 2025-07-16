package main

import (
	"bytes"
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
)

// common
// end common

//end points

// Index EndPoint is the default endpoint that returns a 200 OK status.
func index(_ *Request, respone *Response) {
	respone.Write(200, []byte(Statues[200]))
}

// userAgentReaderEndPoint returns the User-Agent header from the request.
func userAgentReaderEndPoint(request *Request, respone *Response) {
	if userAgent, ok := request.Headers["User-Agent"]; ok {
		respone.Write(200, []byte(userAgent))
		return
	}
	respone.Write(400, []byte(Statues[400]))
}

// echoEndPoint returns the query from the request's endpoint.
func echoEndPoint(request *Request, respone *Response) {
	query := strings.Replace(request.EndPoint, "/echo/", "", 6)
	respone.Write(200, []byte(query))
}

// getFilesEndPoint serves files from the /files/ directory.
func getFilesEndPoint(request *Request, respone *Response) {
	filePath := strings.Replace(request.EndPoint, "/files/", "", 6)
	fullPath := path.Join(AbsPath, filePath)
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
	respone.Write(200, contet)
}

// postFilesEndPoint saves files to the /files/ directory.
func postFilesEndPoint(request *Request, respone *Response) {
	os.Mkdir(AbsPath, os.ModePerm)
	filePath := strings.Replace(request.EndPoint, "/files/", "", 6)
	fullPath := path.Join(AbsPath, filePath)
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		os.Remove(fullPath)
	}
	file, err := os.Create(fullPath)
	if err != nil {
		respone.Write(404, []byte(Statues[404]))
		return
	}

	_, _ = file.Write(request.Body)
	respone.SetHeaders(map[string]string{
		"Content-Length": strconv.Itoa(len(request.Body)),
	})
	respone.Write(201, request.Body)
}

func gzipEndPoint(request *Request, respone *Response) {
	respone.Write(200, []byte{})
}

// end endpoints

//request

type Request struct {
	Method      string
	UserAgent   string
	EndPoint    string
	HTTPVersion string
	Headers     map[string]string
	Body        []byte
}

func NewRequest() *Request {
	return &Request{
		Headers: make(map[string]string),
		Body:    make([]byte, 0),
	}
}

func (req *Request) isValidEndPoint(rgx string) bool {
	regx, err := regexp.Compile(fmt.Sprintf("^%s$", rgx))
	if err != nil {
		return false
	}
	return regx.Match([]byte(req.EndPoint))
}

func Decompress(body []byte) ([]byte, error) {

	reader, err := gzip.NewReader(bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create gzip reader: %w", err)
	}
	defer reader.Close()

	decompressedData, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("error reading decompressed data: %w", err)
	}
	return decompressedData, nil

}

func Compress(body []byte) ([]byte, error) {
	var buff bytes.Buffer
	writer := gzip.NewWriter(&buff)
	defer writer.Close()
	_, err := writer.Write(body)
	if err != nil {
		return []byte{}, err
	}
	return buff.Bytes(), nil
}

func RequestParser(conn net.Conn) (*Request, error) {

	// Accept-Encoding
	// Content-Encoding

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
		request.HTTPVersion = string(requestLine[2])
	}
	if len(requestInfo) > 1 {
		for _, h := range requestInfo[1:] {
			header := bytes.Split(h, []byte(":"))
			if len(header) != 2 || len(bytes.Trim(header[0], " ")) == 0 || len(bytes.Trim(header[1], " ")) == 0 {
				continue
			}
			request.Headers[string(bytes.Trim(header[0], " "))] = string(bytes.Trim(header[1], " "))
		}
		if request.Method == "POST" {
			nullIndex := bytes.IndexByte(requestInfo[len(requestInfo)-1], '\x00')
			if nullIndex == -1 {
				nullIndex = len(requestInfo[len(requestInfo)-1])
			}
			body := requestInfo[len(requestInfo)-1][:nullIndex]

			// if _, ok := request.Headers["Accept-Encoding"]; ok {
			// 	cmps, err := Decompress(body)
			// 	if err != nil {
			// 		return nil, err
			// 	}
			// 	request.Body = cmps
			// 	request.Headers["Content-Encoding"] = "gzip"
			// } else {

			request.Body = body
			fmt.Println(string(request.Body))
			// }
		}
		if _, ok := request.Headers["Accept-Encoding"]; ok {
			request.Headers["Content-Encoding"] = "gzip"
			request.Headers["Accept-Encoding"] = "gzip"
		}

	}
	return request, nil
}

//end request

// response
type Response struct {
	*Request
	conn net.Conn
}

func NewResponse(r *Request, conn net.Conn) *Response {
	return &Response{
		Request: r,
		conn:    conn,
	}
}

func (resp *Response) Write(status int, body []byte) {
	var bd []byte
	if _, ok := resp.Headers["Accept-Encoding"]; ok {
		cmps, err := Decompress(body)
		if err != nil {
			return
		}
		bd = cmps
	} else {
		bd = body
	}
	lenght := len(body)
	responseBody := strings.Builder{}
	responseBody.WriteString(fmt.Sprintf("HTTP/1.1 %d %s\r\n", status, Statues[status]))
	for header, value := range resp.Headers {
		if header == "" || value == "" {
			continue
		}
		responseBody.WriteString(fmt.Sprintf("%s: %s\r\n", header, value))
	}
	if _, ok := resp.Headers["Content-Type"]; !ok {
		responseBody.WriteString("Content-Type: text/plain\r\n")
	}
	if _, ok := resp.Headers["Content-Length"]; !ok {
		responseBody.WriteString(fmt.Sprintf("Content-Length: %d\r\n", lenght))
	}
	if enc, ok := resp.Headers["Accept-Encoding"]; ok {
		resp.Headers["Content-Encoding"] = enc
		resp.Headers["Accept-Encoding"] = enc
	}
	responseBody.WriteString("\r\n")
	responseBody.Write(bd)
	resp.conn.Write([]byte(responseBody.String()))
}

func (resp *Response) SetHeaders(headers map[string]string) {
	for header, value := range headers {
		resp.Headers[header] = value
	}

}

//end response

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

	if request.EndPoint == "/echo/abc" {
		gzipEndPoint(request, response)
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
