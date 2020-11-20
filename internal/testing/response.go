package testing

import (
	"encoding/json"
	"github.com/jpbede/netpalmgo/models"
)

func GetTaskResponseJSON(taskID string, taskStatus models.TaskStatus) []byte {
	enc, _ := json.Marshal(models.Response{
		Status: models.StatusSuccess,
		Data: models.TaskResponse{
			TaskID:     taskID,
			TaskStatus: taskStatus,
		},
	})

	return enc
}
