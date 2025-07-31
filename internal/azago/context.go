// Package azago
package azago

import (
	"context"
	"encoding/json"
	"fmt"
)

type Context struct {
	context.Context
	*Request
	*Response
}

func (_this *Context) BindJSON(obj any) error {
	if err := json.Unmarshal(_this.Body, &obj); err != nil {
		return fmt.Errorf("ctx, could not parse the json : %v", err)
	}
	return nil
}

func (_this *Context) WriteJSON(status int, obj any) error {
	data, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	if err := _this.Write(status, data); err != nil {
		return err
	}
	return nil
}
