package task

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/jpbede/netpalmgo/models"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func getResponse(taskID string, taskStatus models.TaskStatus) models.Response {
	return models.Response{
		Status: models.StatusSuccess,
		Data: models.TaskResponse{
			TaskID:     taskID,
			TaskStatus: taskStatus,
		},
	}
}

func TestClient_GetWithTaskResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// test successful json response
		if req.URL.String() == "/task/123" {
			// Send response to be tested
			resp, err := json.Marshal(getResponse("123", models.TaskStatusFinished))
			assert.NoError(t, err)
			rw.Write(resp)
		}

		// test invalid json response
		if req.URL.String() == "/task/1234" {
			rw.Write([]byte("no json"))
		}

		// test token response
		if req.URL.String() == "/task/12345" {
			http.Error(rw, "{\"detail\":\"Could not validate credentials\"}", http.StatusForbidden)
		}

		// test client error
		if req.URL.String() == "/task/123456" {
			time.Sleep(2 * time.Second)
			http.Error(rw, "{\"detail\":\"Could not validate credentials\"}", http.StatusForbidden)
		}
	}))

	taskCl := New(resty.NewWithClient(server.Client()).SetHostURL(server.URL).SetTimeout(1 * time.Second))

	// test successful response
	getConfigResp := models.TaskResponse{
		TaskID: "123",
	}
	resp, err := taskCl.GetWithTaskResponse(getConfigResp)
	assert.NoError(t, err)
	assert.Equal(t, resp.Data.TaskID, "123")
	assert.Equal(t, resp.Data.TaskStatus, models.TaskStatusFinished)

	// test invalid json response
	getConfigResp = models.TaskResponse{
		TaskID: "1234",
	}
	resp, err = taskCl.GetWithTaskResponse(getConfigResp)
	assert.Error(t, err)
	assert.EqualError(t, err, "invalid character 'o' in literal null (expecting 'u')")

	// test invalid token response
	getConfigResp = models.TaskResponse{
		TaskID: "12345",
	}
	resp, err = taskCl.GetWithTaskResponse(getConfigResp)
	assert.Error(t, err)
	assert.EqualError(t, err, "Could not validate credentials")

	// test client error
	getConfigResp = models.TaskResponse{
		TaskID: "123456",
	}
	resp, err = taskCl.GetWithTaskResponse(getConfigResp)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "context deadline exceeded (Client.Timeout exceeded while awaiting headers)")
}
