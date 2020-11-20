package netpalmgo

import (
	"github.com/jpbede/netpalmgo/apis/getconfig"
	"github.com/jpbede/netpalmgo/apis/setconfig"
	"github.com/jpbede/netpalmgo/apis/task"
	"github.com/jpbede/netpalmgo/internal/transport"
	"net/http"
)

// Client is the main client of netpalmgo
type Client struct {
	transport *transport.Client

	getConfig getconfig.Client
	setConfig setconfig.Client
	task      task.Client
}

// New creates a new Client with APIUrl and APIKey
func New(apiURL, apiKey string) *Client {
	return NewWithClient(apiURL, apiKey, nil)
}

// NewWithClient creates a new Client with APIUrl and APIKey with a given http.Client
func NewWithClient(apiURL, apiKey string, httpClient *http.Client) *Client {
	c := Client{
		transport: transport.NewClient(apiURL, apiKey, httpClient),
	}

	c.getConfig = getconfig.New(c.transport)
	c.setConfig = setconfig.New(c.transport)
	c.task = task.New(c.transport)

	return &c
}

// GetConfig returns the client for the Endpoint /getconfig
func (c *Client) GetConfig() getconfig.Client {
	return c.getConfig
}

// SetConfig returns the client for the Endpoint /setconfig
func (c *Client) SetConfig() setconfig.Client {
	return c.setConfig
}

// Task returns the client for the Endpoint /task
func (c *Client) Task() task.Client {
	return c.task
}
