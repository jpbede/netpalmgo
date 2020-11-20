package task_test

import (
	"context"
	"encoding/json"
	"github.com/jpbede/netpalmgo/apis/task"
	"github.com/jpbede/netpalmgo/internal/transport"
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

func TestClient_WithTaskResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// test successful json response
		if req.URL.String() == "/task/123" {
			// Send response to be tested
			resp, err := json.Marshal(getResponse("123", models.TaskStatusFinished))
			assert.NoError(t, err)
			rw.Write(resp)
		} else if req.URL.String() == "/task/1234" {
			rw.Write([]byte("no json"))
		}
	}))

	httpClient := server.Client()
	httpClient.Timeout = 1 * time.Second
	taskCl := task.New(transport.NewClient(server.URL, "123", httpClient))

	// test successful response
	getConfigResp := models.TaskResponse{
		TaskID: "123",
	}
	resp, err := taskCl.WithTaskResponse(context.Background(), getConfigResp)
	assert.NoError(t, err)
	assert.Equal(t, resp.Data.TaskID, "123")
	assert.Equal(t, resp.Data.TaskStatus, models.TaskStatusFinished)

	// test invalid json response
	getConfigResp = models.TaskResponse{
		TaskID: "1234",
	}
	resp, err = taskCl.WithTaskResponse(context.Background(), getConfigResp)
	assert.Error(t, err)
	assert.EqualError(t, err, "invalid character 'o' in literal null (expecting 'u')")
}
