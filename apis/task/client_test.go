package task_test

import (
	"go.bnck.me/netpalm/apis/task"
	"go.bnck.me/netpalm/internal/transport"
	"testing"
)

func TestNew(t *testing.T) {
	cl := task.New(transport.NewClient("http://localhost", "123", nil))

	if cl == nil {
		t.Error("Got no client")
	}
}
