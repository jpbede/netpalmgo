package task

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/jpbede/netpalmgo/models"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient_GetWithTaskResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test url parameter
		assert.Equal(t, req.URL.String(), "/task/123")

		// Send response to be tested
		resp, err := json.Marshal(models.Response{
			Status: models.StatusSuccess,
			Data: models.TaskResponse{
				TaskID:     "123",
				TaskStatus: models.TaskStatusFinished,
			},
		})
		assert.NoError(t, err)
		rw.Write(resp)
	}))

	taskCl := New(resty.NewWithClient(server.Client()).SetHostURL(server.URL))

	getConfigResp := models.TaskResponse{
		TaskID: "123",
	}
	resp, err := taskCl.GetWithTaskResponse(getConfigResp)
	assert.NoError(t, err)
	assert.Equal(t, resp.Data.TaskID, "123")
	assert.Equal(t, resp.Data.TaskStatus, models.TaskStatusFinished)
}
