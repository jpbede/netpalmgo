package getconfig_test

import (
	"context"
	"encoding/json"
	"github.com/jpbede/netpalmgo/apis/getconfig"
	"github.com/jpbede/netpalmgo/internal/transport"
	"github.com/jpbede/netpalmgo/models"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient_GetWithRequest(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// check path
		assert.Equal(t, req.URL.String(), "/getconfig")

		// check json body
		body, readErr := ioutil.ReadAll(req.Body)
		assert.NoError(t, readErr)
		assert.Equal(t, body, []byte{123, 34, 108, 105, 98, 114, 97, 114, 121, 34, 58, 34, 110, 101, 116, 109, 105, 107, 111, 34, 44, 34, 99, 111, 110, 110, 101, 99, 116, 105, 111, 110, 95, 97, 114, 103, 115, 34, 58, 123, 34, 100, 101, 118, 105, 99, 101, 95, 116, 121, 112, 101, 34, 58, 34, 118, 121, 111, 115, 34, 44, 34, 104, 111, 115, 116, 34, 58, 34, 49, 48, 46, 49, 48, 46, 49, 48, 46, 49, 48, 34, 44, 34, 117, 115, 101, 114, 110, 97, 109, 101, 34, 58, 34, 100, 101, 109, 111, 34, 44, 34, 112, 97, 115, 115, 119, 111, 114, 100, 34, 58, 34, 100, 101, 109, 111, 34, 125, 44, 34, 99, 111, 109, 109, 97, 110, 100, 34, 58, 34, 115, 104, 111, 119, 32, 105, 110, 116, 34, 44, 34, 97, 114, 103, 115, 34, 58, 123, 34, 117, 115, 101, 95, 116, 101, 120, 116, 102, 115, 109, 34, 58, 34, 116, 114, 117, 101, 34, 125, 44, 34, 113, 117, 101, 117, 101, 95, 115, 116, 114, 97, 116, 101, 103, 121, 34, 58, 34, 112, 105, 110, 110, 101, 100, 34, 44, 34, 99, 111, 110, 102, 105, 103, 34, 58, 123, 34, 101, 110, 97, 98, 108, 101, 100, 34, 58, 102, 97, 108, 115, 101, 44, 34, 116, 116, 108, 34, 58, 48, 44, 34, 112, 111, 105, 115, 111, 110, 34, 58, 102, 97, 108, 115, 101, 125, 125, 10})

		// Send response to be tested
		resp, err := json.Marshal(getResponse("123", models.TaskStatusQueued))
		assert.NoError(t, err)
		rw.Write(resp)
	}))

	configCl := getconfig.New(transport.NewClient(server.URL, "123", server.Client()))

	req := models.GetConfigRequest{
		Library: models.LibraryNetmiko,
		ConnectionArgs: models.ConnectionArgs{
			DeviceType: "vyos",
			Host:       "10.10.10.10",
			Username:   "demo",
			Password:   "demo",
		},
		Command:       "show int",
		QueueStrategy: models.QueueStrategyPinned,
	}
	req.UseTextFSM()

	resp, err := configCl.GetWithRequest(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, resp.Status, models.StatusSuccess)
	assert.Equal(t, resp.Data.TaskStatus, models.TaskStatusQueued)
	assert.Equal(t, resp.Data.TaskID, "123")
}
