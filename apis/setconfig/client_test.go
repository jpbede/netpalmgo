package setconfig

import (
	"github.com/go-resty/resty/v2"
	"testing"
)

func TestNew(t *testing.T) {
	cl := New(resty.New())

	if cl == nil {
		t.Error("Got no client")
	}
}
