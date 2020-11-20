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
		assert.Equal(t, body, []byte{0x7b, 0x22, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x22, 0x3a, 0x22, 0x6e, 0x65, 0x74, 0x6d, 0x69, 0x6b, 0x6f, 0x22, 0x2c, 0x22, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x72, 0x67, 0x73, 0x22, 0x3a, 0x7b, 0x22, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x22, 0x76, 0x79, 0x6f, 0x73, 0x22, 0x2c, 0x22, 0x68, 0x6f, 0x73, 0x74, 0x22, 0x3a, 0x22, 0x31, 0x30, 0x2e, 0x31, 0x30, 0x2e, 0x31, 0x30, 0x2e, 0x31, 0x30, 0x22, 0x2c, 0x22, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3a, 0x22, 0x64, 0x65, 0x6d, 0x6f, 0x22, 0x2c, 0x22, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x3a, 0x22, 0x64, 0x65, 0x6d, 0x6f, 0x22, 0x7d, 0x2c, 0x22, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x22, 0x3a, 0x5b, 0x22, 0x73, 0x68, 0x6f, 0x77, 0x20, 0x69, 0x6e, 0x74, 0x22, 0x5d, 0x2c, 0x22, 0x71, 0x75, 0x65, 0x75, 0x65, 0x5f, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x22, 0x3a, 0x22, 0x66, 0x69, 0x66, 0x6f, 0x22, 0x2c, 0x22, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x3a, 0x7b, 0x22, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x22, 0x3a, 0x66, 0x61, 0x6c, 0x73, 0x65, 0x2c, 0x22, 0x74, 0x74, 0x6c, 0x22, 0x3a, 0x30, 0x2c, 0x22, 0x70, 0x6f, 0x69, 0x73, 0x6f, 0x6e, 0x22, 0x3a, 0x66, 0x61, 0x6c, 0x73, 0x65, 0x7d, 0x7d, 0xa})

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
