// Package zenith
package zenith

import (
	"context"

	"github.com/maruki00/zenithgo/internal/http/request"
	"github.com/maruki00/zenithgo/internal/http/response"
)

type Context struct {
	context.Context
	*request.Request
	response.Response
}
