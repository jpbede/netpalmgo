package netpalmgo

import (
	"net/http"
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

	if cl.resty == nil {
		t.Error("Resty client is empty")
	}
}

func TestNewWithClient(t *testing.T) {
	cl := NewWithClient(os.Getenv("NETPALM_APIURL"), os.Getenv("NETPALM_APIKEY"), &http.Client{})

	if cl.resty == nil {
		t.Error("Resty client is empty")
	}
}

func TestClient_GetConfig(t *testing.T) {
	cl := GetClient()
	cfgAPI := cl.GetConfig()

	if cfgAPI == nil {
		t.Error("Failed to get 'GetConfig' endpoint")
	}
}
