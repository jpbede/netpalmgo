package models

import (
	"encoding/json"
)

// TaskResults represents the results of the task
type TaskResults map[string]interface{}

// MapTo marshals task result to given struct
func (tr *TaskResults) MapTo(destStruct interface{}) error {
	enc, err := json.Marshal(tr)
	if err != nil {
		return err
	}
	err = json.Unmarshal(enc, destStruct)
	if err != nil {
		return err
	}
	return nil
}

// TaskError represents a error
type TaskError string

// TaskResponse represents informations about the task
type TaskResponse struct {
	TaskID       string       `json:"task_id"`
	CreatedOn    string       `json:"created_on"`
	TaskQueue    string       `json:"task_queue"`
	TaskMetaData TaskMetaData `json:"task_meta"`
	TaskStatus   TaskStatus   `json:"task_status"`
	TaskResult   TaskResults  `json:"task_result"`
	TaskErrors   []TaskError  `json:"task_errors"`
}
