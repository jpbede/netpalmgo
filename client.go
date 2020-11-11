package netpalmgo

import (
	"github.com/go-resty/resty/v2"
	"github.com/jpbede/netpalmgo/apis/getconfig"
	"net/http"
)

type Client struct {
	resty *resty.Client

	getConfig getconfig.Client
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

func New(apiUrl, apiKey string) *Client {
	c := Client{}
	c.resty = newResty(apiUrl, apiKey, nil)
	return &c
}

func NewWithClient(apiUrl, apiKey string, httpClient *http.Client) *Client {
	c := Client{}
	c.resty = newResty(apiUrl, apiKey, httpClient)
	return &c
}

func (c *Client) GetConfig() getconfig.Client {
	if c.getConfig == nil {
		c.getConfig = getconfig.New(c.resty)
	}
	return c.getConfig
}
