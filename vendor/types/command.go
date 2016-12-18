package types

type CommandType string

const (
	CommandCreate CommandType = "create"
	CommandStart  CommandType = "start"
	CommandStop   CommandType = "stop"
	CommandState  CommandType = "state"
)
