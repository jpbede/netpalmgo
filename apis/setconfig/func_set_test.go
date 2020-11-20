package setconfig_test

import (
	"context"
	"github.com/jpbede/netpalmgo/apis/setconfig"
	"github.com/jpbede/netpalmgo/internal/transport"
	"github.com/jpbede/netpalmgo/models"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient_Set(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// check if right path
		assert.Equal(t, req.URL.String(), "/setconfig")

		// check request body
		body, err := ioutil.ReadAll(req.Body)
		assert.NoError(t, err)
		assert.Equal(t, []byte{0x7b, 0x22, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x22, 0x3a, 0x22, 0x6e, 0x61, 0x70, 0x61, 0x6c, 0x6d, 0x22, 0x2c, 0x22, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x72, 0x67, 0x73, 0x22, 0x3a, 0x7b, 0x22, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x22, 0x76, 0x79, 0x6f, 0x73, 0x22, 0x2c, 0x22, 0x68, 0x6f, 0x73, 0x74, 0x22, 0x3a, 0x22, 0x31, 0x30, 0x2e, 0x31, 0x30, 0x2e, 0x31, 0x30, 0x2e, 0x31, 0x30, 0x22, 0x2c, 0x22, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3a, 0x22, 0x64, 0x65, 0x6d, 0x6f, 0x22, 0x2c, 0x22, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x3a, 0x22, 0x64, 0x65, 0x6d, 0x6f, 0x22, 0x7d, 0x2c, 0x22, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x3a, 0x5b, 0x22, 0x73, 0x65, 0x74, 0x20, 0x69, 0x6e, 0x74, 0x20, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x20, 0x74, 0x75, 0x6e, 0x30, 0x20, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2d, 0x69, 0x70, 0x20, 0x31, 0x32, 0x33, 0x2e, 0x31, 0x32, 0x33, 0x2e, 0x31, 0x32, 0x33, 0x2e, 0x31, 0x32, 0x33, 0x22, 0x5d, 0x2c, 0x22, 0x71, 0x75, 0x65, 0x75, 0x65, 0x5f, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x22, 0x3a, 0x22, 0x66, 0x69, 0x66, 0x6f, 0x22, 0x7d, 0xa}, body)

		rw.Write([]byte("{\"status\":\"success\",\"data\":{\"task_id\":\"b88975a9-d460-461a-89f8-f5f9c3dc8701\",\"created_on\":\"2020-11-20 09:56:22.152880\",\"task_queue\":\"fifo\",\"task_meta\":{\"enqueued_at\":\"2020-11-20 09:56:22.153401\",\"started_at\":null,\"ended_at\":null,\"enqueued_elapsed_seconds\":\"0\",\"total_elapsed_seconds\":\"0\"},\"task_status\":\"queued\",\"task_result\":null,\"task_errors\":[]}}"))
	}))

	setConfigClient := setconfig.New(transport.NewClient(server.URL, "123", server.Client()))

	setConfigReq := models.SetConfigRequest{
		Library: models.LibraryNapalm,
		ConnectionArgs: models.ConnectionArgs{
			DeviceType: "vyos",
			Host:       "10.10.10.10",
			Username:   "demo",
			Password:   "demo",
		},
		Config:        []string{"set int tunnel tun0 remote-ip 123.123.123.123"},
		QueueStrategy: models.QueueStrategyFIFO,
	}

	resp, err := setConfigClient.Set(context.Background(), false, setConfigReq)
	assert.NoError(t, err)
	assert.Equal(t, models.TaskStatusQueued, resp.Data.TaskStatus)
	assert.Equal(t, "b88975a9-d460-461a-89f8-f5f9c3dc8701", resp.Data.TaskID)
}

func TestClient_SetInvalidJson(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte("no json"))
	}))

	setConfigClient := setconfig.New(transport.NewClient(server.URL, "123", server.Client()))

	setConfigReq := models.SetConfigRequest{
		Library: models.LibraryNapalm,
		ConnectionArgs: models.ConnectionArgs{
			DeviceType: "vyos",
			Host:       "10.10.10.10",
			Username:   "demo",
			Password:   "demo",
		},
		Config:        []string{"set int tunnel tun0 remote-ip 123.123.123.123"},
		QueueStrategy: models.QueueStrategyFIFO,
	}

	_, err := setConfigClient.Set(context.Background(), false, setConfigReq)
	assert.Error(t, err)
	assert.EqualError(t, err, "invalid character 'o' in literal null (expecting 'u')")
}
