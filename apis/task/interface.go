package task

import "github.com/jpbede/netpalmgo/models"

type Client interface {
	GetWithTaskResponse(taskResponse models.TaskResponse) (*models.Response, error)
}
