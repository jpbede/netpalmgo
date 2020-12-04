package models

// TaskStatus represents the status of a task
type TaskStatus string

const (
	// TaskStatusQueued task execution is queued
	TaskStatusQueued TaskStatus = "queued"
	// TaskStatusFinished task execution is finished
	TaskStatusFinished TaskStatus = "finished"
	// TaskStatusFailed task execution failed
	TaskStatusFailed TaskStatus = "failed"
	// TaskStatusStarted task execution started
	TaskStatusStarted TaskStatus = "started"
	// TaskStatusDeferred task execution deferred
	TaskStatusDeferred TaskStatus = "deferred"
	// TaskStatusScheduled task execution scheduled
	TaskStatusScheduled TaskStatus = "scheduled"
)
