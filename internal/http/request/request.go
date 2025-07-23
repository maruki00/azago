//Package request
package request

import (
	"bytes"
	"errors"
	"net"

	"github.com/maruki00/zenithgo/internal/common"
	gzipPkg "github.com/maruki00/zenithgo/pkg/gzip"
	readerPkg "github.com/maruki00/zenithgo/pkg/reader"
)

type Request struct {
	Method      string
	UserAgent   string
	EndPoint    string
	HTTPVersion string
	Headers     map[string]string
	Body        []byte
}

func NewRequest(conn net.Conn) *Request {
	req := &Request{
		Headers: make(map[string]string),
		Body:    make([]byte, 0),
	}
	if err := req.RequestParser(conn); err != nil {
		return nil
	}
	return req
}

func (_this *Request) RequestParser(conn net.Conn) error {
	requestBuff, err := readerPkg.ReadUntil(conn, []byte("\r\n\r\n"))
	if err != nil {
		return err
	}

	requestInfo := bytes.Split(requestBuff, []byte(common.CRLF))
	if len(requestInfo) == 0 {
		return errors.New("request is empty")
	}
	requestLine := bytes.Split(requestInfo[0], []byte(" "))
	if len(requestLine) == 0 {
		return errors.New("invalide request line")
	}

	rLineLenght := len(requestLine)
	if rLineLenght >= 1 {
		_this.Method = string(requestLine[0])
	}
	if rLineLenght >= 2 {
		_this.EndPoint = string(requestLine[1])
	}
	if rLineLenght >= 3 {
		_this.HTTPVersion = string(requestLine[2])
	}
	var header [][]byte
	for _, h := range requestInfo[1:] {
		header = bytes.Split(h, []byte(":"))
		if len(header) != 2 || len(bytes.Trim(header[0], " ")) == 0 || len(bytes.Trim(header[1], " ")) == 0 {
			continue
		}
		_this.Headers[string(bytes.Trim(header[0], " "))] = string(bytes.Trim(header[1], " "))
	}
	if _this.Method == "POST" {
		nullIndex := bytes.IndexByte(requestInfo[len(requestInfo)-1], '\x00')
		if nullIndex == -1 {
			nullIndex = len(requestInfo[len(requestInfo)-1])
		}
		body := requestInfo[len(requestInfo)-1][:nullIndex]

		if _, ok := _this.Headers["Accept-Encoding"]; ok {
			body, err := gzipPkg.Decompress(body)
			if err != nil {
				return err
			}
			_this.Body = body
			_this.Headers["Content-Encoding"] = "gzip"
		} else {
			_this.Body = body
		}
	}
	return nil
}
