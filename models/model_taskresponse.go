package models

type TaskResult []string
type TaskResults map[string]TaskResult
type TaskError map[string]string

type TaskResponse struct {
	TaskID       string       `json:"task_id"`
	CreatedOn    string       `json:"created_on"`
	TaskQueue    string       `json:"task_queue"`
	TaskMetaData TaskMetaData `json:"task_meta"`
	TaskStatus   TaskStatus   `json:"task_status"`
	TaskResult   TaskResults  `json:"task_result"`
	TaskErrors   []TaskError  `json:"task_errors"`
}
