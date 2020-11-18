package transport

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

// Client is the http transport client for netpalm. It handles the authentication
type Client struct {
	baseURL    string
	httpClient *http.Client
	apiKey     string
}

// NewClient returns a new Transport HTTP client
func NewClient(baseURL, apiKey string, hc *http.Client) *Client {
	if hc == nil {
		hc = &http.Client{}
	}

	c := Client{
		baseURL:    baseURL,
		apiKey:     apiKey,
		httpClient: hc,
	}
	return &c
}

// Get executes a GET request
func (c *Client) Get(ctx context.Context, path string, out interface{}, opts ...RequestOption) error {
	return c.doRequest(ctx, http.MethodGet, path, out, opts...)
}

// Post executes a POST request
func (c *Client) Post(ctx context.Context, path string, out interface{}, opts ...RequestOption) error {
	return c.doRequest(ctx, http.MethodPost, path, out, opts...)
}

// Put executes a PUT request
func (c *Client) Put(ctx context.Context, path string, out interface{}, opts ...RequestOption) error {
	return c.doRequest(ctx, http.MethodPut, path, out, opts...)
}

// Patch executes a PATCH request
func (c *Client) Patch(ctx context.Context, path string, out interface{}, opts ...RequestOption) error {
	return c.doRequest(ctx, http.MethodPatch, path, out, opts...)
}

// Delete executes a DELETE request
func (c *Client) Delete(ctx context.Context, path string, out interface{}, opts ...RequestOption) error {
	return c.doRequest(ctx, http.MethodPatch, path, out, opts...)
}

// doRequest does the actual request
func (c *Client) doRequest(ctx context.Context, method string, path string, out interface{}, opts ...RequestOption) error {
	// create a new request
	path = strings.TrimPrefix(path, "/")
	req, err := http.NewRequest(method, c.baseURL+"/"+path, nil)
	if err != nil {
		return err
	}
	// add api key
	req.Header.Set("x-api-key", c.apiKey)
	// run options
	for i := range opts {
		if err := opts[i](req); err != nil {
			return err
		}
	}
	req = req.WithContext(ctx)

	// run request
	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	httpErr := CheckForHTTPError(res)
	if httpErr != nil {
		return httpErr
	}

	if out != nil {
		if w, ok := out.(io.Writer); ok {
			_, err := io.Copy(w, res.Body)
			return err
		}

		if err := json.NewDecoder(res.Body).Decode(out); err != nil {
			return err
		}
	}

	return nil
}
