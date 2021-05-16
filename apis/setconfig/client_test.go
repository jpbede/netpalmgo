package setconfig_test

import (
	"go.bnck.me/netpalm/apis/setconfig"
	"go.bnck.me/netpalm/internal/transport"
	"testing"
)

func TestNew(t *testing.T) {
	cl := setconfig.New(transport.NewClient("http://localhost", "123", nil))

	if cl == nil {
		t.Error("Got no client")
	}
}
