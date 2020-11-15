package models

// TaskResults represents the results of the task
type TaskResults map[string]interface{}

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
