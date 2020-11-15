package models

// QueueStrategy how a task would be handled
type QueueStrategy string

const (
	QueueStrategyFIFO   QueueStrategy = "fifo"
	QueueStrategyPinned QueueStrategy = "pinned"
)
