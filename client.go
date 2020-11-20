package netpalmgo

import (
	"context"
	"github.com/jpbede/netpalmgo/apis/getconfig"
	"github.com/jpbede/netpalmgo/apis/setconfig"
	"github.com/jpbede/netpalmgo/apis/task"
	"github.com/jpbede/netpalmgo/internal/transport"
	"github.com/jpbede/netpalmgo/models"
	"net/http"
	"sync"
	"time"
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

// WaitForResult waits blocking for the result to finish (also fail).
func (c *Client) WaitForResult(ctx context.Context, response *models.Response) (*models.Response, error) {
	wg := sync.WaitGroup{}
	resp := response
	var err error

	// wait for task to finish in other goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()

		// sleep for 2 sec if task is not finished or failed and then check status again
		for resp.Data.TaskStatus != models.TaskStatusFinished && resp.Data.TaskStatus != models.TaskStatusFailed {
			resp, err = c.Task().WithTaskResponse(ctx, resp.Data)
			if err != nil { // exit when error occurs
				return
			}
			time.Sleep(2 * time.Second)
		}
	}()
	wg.Wait()

	if err != nil {
		return nil, err
	}

	return resp, err
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
