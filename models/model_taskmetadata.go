package models

type TaskMetaData struct {
	EnqueuedAt             string `json:"enqueued_at"`
	StartedAt              string `json:"started_at"`
	EndedAt                string `json:"ended_at"`
	EnqueuedElapsedSeconds string `json:"enqueued_elapsed_seconds"`
	TotalElapsedSeconds    string `json:"total_elapsed_seconds"`
}
