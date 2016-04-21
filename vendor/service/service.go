package service

import "types"

type AgentService interface {
	State() types.AgentState
	Start() error
	Stop() error
	Name() string
	Operate(id, operate string, data interface{}) (error, interface{})
}

type ServiceMiddleware func(AgentService) AgentService
