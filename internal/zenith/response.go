// Package zenith
package zenith

import (
	"fmt"
	"net"
	"strings"

	"github.com/maruki00/zenithgo/internal/common"
)

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

func (_this *Response) Write(status int, body []byte) error{
	lenght := len(body)
	responseBody := strings.Builder{}
	responseBody.WriteString(fmt.Sprintf("HTTP/1.1 %d %s\r\n", status, common.Statues[status]))
	for header, value := range _this.Headers {
		if header == "" || value == "" {
			continue
		}
		responseBody.WriteString(fmt.Sprintf("%s: %s\r\n", header, value))
	}
	if _, ok := _this.Headers["Content-Type"]; !ok {
		responseBody.WriteString("Content-Type: text/plain\r\n")
	}
	if _, ok := _this.Headers["Content-Length"]; !ok {
		responseBody.WriteString(fmt.Sprintf("Content-Length: %d\r\n", lenght))
	}
	responseBody.WriteString("\r\n")
	responseBody.Write(body)
	_,err := _this.conn.Write([]byte(responseBody.String()))
	return err
}

func (_this *Response) SetHeaders(headers map[string]string) {
	for header, value := range headers {
		_this.Headers[header] = value
	}
}
