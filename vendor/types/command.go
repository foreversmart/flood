package types

type CommandType string

const (
	CommandStart CommandType = "start"
	CommandStop  CommandType = "stop"
	CommandState CommandType = "state"
)
