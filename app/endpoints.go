package main

import (
	"os"
	"path"
	"strconv"
	"strings"
)

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
