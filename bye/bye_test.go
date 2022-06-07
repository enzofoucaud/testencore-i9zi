package bye

import (
	"context"
	"testing"
)

func TestBye(t *testing.T) {
	resp, err := Bye(context.Background(), "Jane Doe")
	if err != nil {
		t.Fatal(err)
	}
	want := "Bye, Jane Doe!"
	if got := resp.Message; got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
