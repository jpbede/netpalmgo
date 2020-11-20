package getconfig_test

import (
	"context"
	"github.com/jpbede/netpalmgo/apis/getconfig"
	testing2 "github.com/jpbede/netpalmgo/internal/testing"
	"github.com/jpbede/netpalmgo/internal/transport"
	"github.com/jpbede/netpalmgo/models"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient_Get(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// check path
		assert.Equal(t, req.URL.String(), "/getconfig")

		// check json body
		body, readErr := ioutil.ReadAll(req.Body)
		assert.NoError(t, readErr)
		assert.Equal(t, testing2.GetExpectedTaskRequest(), body)

		// Send response to be tested
		rw.Write(testing2.GetTaskResponseJSON("123", models.TaskStatusQueued))
	}))

	configCl := getconfig.New(transport.NewClient(server.URL, "123", server.Client()))

	resp, err := configCl.WithCommand(context.Background(), "show int", models.LibraryNetmiko, testing2.GetConnectionArgs())
	assert.NoError(t, err)
	assert.Equal(t, resp.Status, models.StatusSuccess)
	assert.Equal(t, resp.Data.TaskStatus, models.TaskStatusQueued)
	assert.Equal(t, resp.Data.TaskID, "123")
}
