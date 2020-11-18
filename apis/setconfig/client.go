package setconfig

import (
	"github.com/jpbede/netpalmgo/internal/transport"
)

type client struct {
	transport *transport.Client
}

// New creates a new Client for the setconfig Endpoint
func New(transport *transport.Client) Client {
	return &client{
		transport: transport,
	}
}
