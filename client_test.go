package netpalmgo

import (
	"context"
	testing2 "github.com/jpbede/netpalmgo/internal/testing"
	"github.com/jpbede/netpalmgo/models"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var client *Client

func GetClient() *Client {
	if client != nil {
		return client
	}
	client = New(os.Getenv("NETPALM_APIURL"), os.Getenv("NETPALM_APIKEY"))
	return client
}

func TestNew(t *testing.T) {
	cl := New(os.Getenv("NETPALM_APIURL"), os.Getenv("NETPALM_APIKEY"))

	if cl.transport == nil {
		t.Error("Resty client is empty")
	}
}

func TestNewWithClient(t *testing.T) {
	cl := NewWithClient(os.Getenv("NETPALM_APIURL"), os.Getenv("NETPALM_APIKEY"), &http.Client{})

	if cl.transport == nil {
		t.Error("Resty client is empty")
	}
}

func TestWaitForResult(t *testing.T) {
	reqcounter := 0

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Send response to be tested
		if reqcounter > 2 {
			rw.Write(testing2.GetTaskResponseJSON("123", models.TaskStatusFinished))
		} else {
			rw.Write(testing2.GetTaskResponseJSON("123", models.TaskStatusQueued))
		}

		reqcounter++
	}))
	cl := NewWithClient(server.URL, "123", server.Client())

	resp, err := cl.GetConfig().WithCommand(context.Background(), "show int", models.LibraryNetmiko, testing2.GetConnectionArgs())
	assert.NoError(t, err)

	resp, err = cl.WaitForResult(context.Background(), resp)
	assert.NoError(t, err)
	assert.Equal(t, models.TaskStatusFinished, resp.Data.TaskStatus)
}

func TestClient_GetConfig(t *testing.T) {
	cl := GetClient()
	cfgAPI := cl.GetConfig()

	if cfgAPI == nil {
		t.Error("Failed to get 'GetConfig' endpoint")
	}
}

func TestClient_SetConfig(t *testing.T) {
	cl := GetClient()
	cfgAPI := cl.SetConfig()

	if cfgAPI == nil {
		t.Error("Failed to get 'SetConfig' endpoint")
	}
}

func TestClient_Task(t *testing.T) {
	cl := GetClient()
	taskAPI := cl.Task()

	if taskAPI == nil {
		t.Error("Failed to get 'Task' endpoint")
	}
}
