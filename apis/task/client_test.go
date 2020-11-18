package task_test

import (
	"github.com/jpbede/netpalmgo/apis/task"
	"github.com/jpbede/netpalmgo/internal/transport"
	"testing"
)

func TestNew(t *testing.T) {
	cl := task.New(transport.NewClient("http://localhost", "123", nil))

	if cl == nil {
		t.Error("Got no client")
	}
}
