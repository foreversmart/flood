package service

type AgentService interface {
	Operate(id, operate string, data interface{}) (error, interface{})
}

type ServiceMiddleware func(AgentService) AgentService
