package bye

import (
	"context"
)

//encore:api public path=/bye/:name
func Bye(ctx context.Context, name string) (*Response, error) {
	msg := "Bye, " + name + "!"
	return &Response{Message: msg}, nil
}

type Response struct {
	Message string
}
