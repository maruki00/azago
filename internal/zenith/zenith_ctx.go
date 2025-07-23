// Package zenith
package zenith

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
)

type Context struct {
	context.Context
	*Request
	*Response
}


func (_this *Context)BindJson(obj interface{}) error {
	if err := json.Unmarshal(_this.Body, &obj); err != nil {
		return fmt.Errorf("could not parse the json : %v", err)
	}
	return nil
}

func (_this *Context) WriteJson(status int, obj interface{}) error {
	data , err := json.Marshal(obj)
	if err != nil {
		return err
	}
	if _,err :=  _this.Write(status, data); err != nil {
		return err
	}
	return nil
}












