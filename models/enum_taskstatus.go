package models

type TaskStatus string

const (
	TaskStatusQueued    TaskStatus = "queued"
	TaskStatusFinished  TaskStatus = "finished"
	TaskStatusFailed    TaskStatus = "failed"
	TaskStatusStarted   TaskStatus = "started"
	TaskStatusDeferred  TaskStatus = "deferred"
	TaskStatusScheduled TaskStatus = "scheduled"
)
