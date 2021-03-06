package hello

import (
	"context"
	"testing"
)

func TestBye(t *testing.T) {
	resp, err := Hello(context.Background(), "Jane Doe")
	if err != nil {
		t.Fatal(err)
	}
	want := "Hello, Jane Doe!"
	if got := resp.Message; got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
