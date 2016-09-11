package service

import "types"

type AgentService interface {
	Operate(id, operate types.CommandType, data interface{}) (error, interface{})
}

type ServiceMiddleware func(AgentService) AgentService
