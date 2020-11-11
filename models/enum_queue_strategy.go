package models

type QueueStrategy string

const (
	QueueStrategyFIFO   QueueStrategy = "fifo"
	QueueStrategyPinned QueueStrategy = "pinned"
)
