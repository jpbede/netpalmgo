package setconfig

import "github.com/go-resty/resty/v2"

type client struct {
	transport *resty.Client
}

func New(transport *resty.Client) Client {
	return client{
		transport: transport,
	}
}
