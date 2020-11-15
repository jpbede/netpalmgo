package netpalmgo

import (
	"github.com/jpbede/netpalmgo/apis/getconfig"
	"github.com/jpbede/netpalmgo/apis/task"
	"net/http"
)

// Client is the main client of netpalmgo
type Client struct {
	resty *resty.Client

	getConfig getconfig.Client
	task      task.Client
}

func newResty(apiUrl, apiKey string, httpClient *http.Client) *resty.Client {
	var r *resty.Client
	if httpClient != nil {
		r = resty.NewWithClient(httpClient)
	} else {
		r = resty.New()
	}
	r.SetHostURL(apiUrl)
	r.SetHeader("x-api-key", apiKey)
	return r
}

// New creates a new Client with APIUrl and APIKey
func New(apiUrl, apiKey string) *Client {
	c := Client{}
	c.resty = newResty(apiUrl, apiKey, nil)
	return &c
}

// New creates a new Client with APIUrl and APIKey with a given http.Client
func NewWithClient(apiUrl, apiKey string, httpClient *http.Client) *Client {
	c := Client{}
	c.resty = newResty(apiUrl, apiKey, httpClient)
	return &c
}

// GetConfig returns the client for the Endpoint /getconfig
func (c *Client) GetConfig() getconfig.Client {
	if c.getConfig == nil {
		c.getConfig = getconfig.New(c.resty)
	}
	return c.getConfig
}

// Task returns the client for the Endpoint /task
func (c *Client) Task() task.Client {
	if c.task == nil {
		c.task = task.New(c.resty)
	}
	return c.task
}
