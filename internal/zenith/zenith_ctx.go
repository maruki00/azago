// Package zenith
package zenith

import (
	"context"
)

type Context struct {
	context.Context
	*Request
	*Response
}
