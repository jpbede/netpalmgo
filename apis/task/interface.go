package task

import "github.com/jpbede/netpalmgo/models"

// Client represents the functions implemented by this API
type Client interface {
	GetWithTaskResponse(taskResponse models.TaskResponse) (*models.Response, error)
}
