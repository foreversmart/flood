package types

const (
	AgentStateMin AgentState = iota
	AgentStateStop
	AgentStateRunning
	AgentStateMax
)

type AgentState int

func (as AgentState) IsValid() bool {
	return as > AgentStateMin && as < AgentStateMax
}
