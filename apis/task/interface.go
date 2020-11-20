package task

import (
	"context"
	"github.com/jpbede/netpalmgo/models"
)

// Client represents the functions implemented by this API
type Client interface {
	WithTaskResponse(ctx context.Context, taskResponse models.TaskResponse) (*models.Response, error)
}
