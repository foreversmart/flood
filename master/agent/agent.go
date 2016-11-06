package agent

import (
	"errors"
	"fmt"
	"types"
)

type DefaultAgent struct {
	state types.AgentState
	name  string
}

func (as DefaultAgent) State() types.AgentState {
	return as.state
}

func (as DefaultAgent) Name() string {
	return as.name
}

func (as DefaultAgent) Start() error {
	fmt.Println(as.name + "start")
	return nil
}

func (as DefaultAgent) Stop() error {
	fmt.Println(as.name + "stop")
	return nil
}

func (as DefaultAgent) Operate(id string, operate types.CommandType, data interface{}) (error, interface{}) {
	fmt.Println(as.name + "operate")
	switch operate {
	case "start":
		return as.Start(), nil
	case "stop":
		return as.Stop(), nil
	case "name":
		return nil, as.Name()
	case "state":
		return nil, as.State()
	default:
		return errors.New("cant surpport operate type"), nil
	}
}
