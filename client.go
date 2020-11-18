package netpalmgo

import (
	"github.com/jpbede/netpalmgo/apis/getconfig"
	"github.com/jpbede/netpalmgo/apis/task"
	"github.com/jpbede/netpalmgo/internal/transport"
	"net/http"
)

// Client is the main client of netpalmgo
type Client struct {
	transport *transport.Client

	getConfig getconfig.Client
	task      task.Client
}

func newTransport(apiURL, apiKey string, httpClient *http.Client) *transport.Client {
	return transport.NewClient(apiURL, apiKey, httpClient)
}

// New creates a new Client with APIUrl and APIKey
func New(apiURL, apiKey string) *Client {
	c := Client{}
	c.transport = newTransport(apiURL, apiKey, nil)
	return &c
}

// NewWithClient creates a new Client with APIUrl and APIKey with a given http.Client
func NewWithClient(apiURL, apiKey string, httpClient *http.Client) *Client {
	c := Client{}
	c.transport = newTransport(apiURL, apiKey, httpClient)
	return &c
}

// GetConfig returns the client for the Endpoint /getconfig
func (c *Client) GetConfig() getconfig.Client {
	if c.getConfig == nil {
		c.getConfig = getconfig.New(c.transport)
	}
	return c.getConfig
}

// Task returns the client for the Endpoint /task
func (c *Client) Task() task.Client {
	if c.task == nil {
		c.task = task.New(c.transport)
	}
	return c.task
}
