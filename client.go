package netpalmgo

import (
	"github.com/go-resty/resty/v2"
	"github.com/jpbede/netpalmgo/apis/getconfig"
	"net/http"
)

type Client struct {
	resty *resty.Client

	getconfig getconfig.Client
}

func New(apiUrl, apiKey string) Client {
	c := Client{}
	c.resty = resty.New()
	c.resty.SetHostURL(apiUrl)
	c.resty.SetHeader("x-api-key", apiKey)

	c.getconfig = getconfig.New(c.resty)

	return c
}

func NewWithClient(apiUrl, apiKey string, httpClient *http.Client) Client {
	c := Client{}
	c.resty = resty.NewWithClient(httpClient)
	c.resty.SetHostURL(apiUrl)
	c.resty.SetHeader("x-api-key", apiKey)

	c.getconfig = getconfig.New(c.resty)

	return c
}

func (c *Client) GetConfig() getconfig.Client {
	return c.getconfig
}
