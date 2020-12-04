package models

// QueueStrategy how a task would be handled
type QueueStrategy string

const (
	// QueueStrategyFIFO queue strategy first in first out
	QueueStrategyFIFO QueueStrategy = "fifo"
	// QueueStrategyPinned queue strategy with a fixed worker
	QueueStrategyPinned QueueStrategy = "pinned"
)
