package setconfig_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go.bnck.me/netpalm/apis/setconfig"
	"go.bnck.me/netpalm/internal/transport"
	"go.bnck.me/netpalm/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDryRun(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// check if dry-run path
		assert.Equal(t, req.URL.String(), "/setconfig/dry-run")

		rw.Write([]byte("{}"))
	}))

	setConfigClient := setconfig.New(transport.NewClient(server.URL, "123", server.Client()))

	setConfigReq := models.SetConfigRequest{}
	_, err := setConfigClient.DryRun(context.Background(), setConfigReq)
	assert.NoError(t, err)
}
