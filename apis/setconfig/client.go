package setconfig

import "github.com/go-resty/resty/v2"

type client struct {
	transport *resty.Client
}

// New creates a new Client for the setconfig Endpoint
func New(transport *resty.Client) Client {
	return &client{
		transport: transport,
	}
}
