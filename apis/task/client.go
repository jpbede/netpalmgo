package task

import "github.com/go-resty/resty/v2"

type client struct {
	transport *resty.Client
}

// New creates a new Client for the task Endpoint
func New(transport *resty.Client) Client {
	return &client{
		transport: transport,
	}
}
