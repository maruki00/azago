package request

import (
	"bytes"
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"net"
	"regexp"

	"github.com/eapache/go-resiliency/breaker"
	"github.com/maruki00/zenithgo/internal/common"
	readerPkg "github.com/maruki00/zenithgo/internal/pkg/reader"
)

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

func (_this*Request) RequestParser(conn net.Conn) (*Request, error) {	
	request := NewRequest()
	requestBuff := readerPkg.Read(conn)
	
	requestInfo := bytes.Split(requestBuff, []byte(common.CRLF))
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

			if _, ok := request.Headers["Accept-Encoding"]; ok {
				cmps, err := Decompress(body)
				if err != nil {
					return nil, err
				}
				request.Body = cmps
				request.Headers["Content-Encoding"] = "gzip"
			} else {

				request.Body = body 				
				fmt.Println(string(request.Body))
			}
		}
	}
	return request, nil
}
