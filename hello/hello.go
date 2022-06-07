package hello

import (
	"context"
)

//encore:api public path=/hello/:name
func Hello(ctx context.Context, name string) (*Response, error) {
	msg := "Hello, " + name + "!"
	return &Response{Message: msg}, nil
}

type Response struct {
	Message string
}
