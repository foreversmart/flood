package agent

import (
	"errors"
	"flood/slave/executor"
	"flood/slave/task_manager"
	"fmt"
	"types"
)

type DefaultAgent struct {
	state       types.AgentState
	name        string
	TaskManager *task_manager.TaskManager
}

func NewAgent() DefaultAgent {
	return DefaultAgent{
		TaskManager: task_manager.NewTaskManager(),
	}
}

func (as DefaultAgent) State() types.AgentState {
	return as.state
}

func (as DefaultAgent) Start(task *types.Task) error {
	as.TaskManager.Add(task)
	executor.Executor.Run(task)
	fmt.Println(as.name + "start")
	return nil
}

func (as DefaultAgent) Stop() error {
	fmt.Println(as.name + "stop")
	return nil
}

func (as DefaultAgent) Operate(id, operate types.CommandType, data interface{}) (error, interface{}) {
	fmt.Println(as.name + "operate")
	switch operate {
	case types.CommandStart:
		return as.Start(data.(*types.Task)), nil
	case types.CommandStop:
		return as.Stop(), nil
	case types.CommandState:
		return nil, as.State()
	default:
		return errors.New("cant surpport operate type"), nil
	}
}
