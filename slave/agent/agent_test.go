package agent

import (
	"testing"
	"types"

	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

var (
	testAgent DefaultAgent
)

func TestDefaultAgentCommand(t *testing.T) {
	testAgent = NewAgent()
	t.Run("test", TestDefaultAgent_Create)
}

func TestDefaultAgent_Create(t *testing.T) {
	assertion := assert.New(t)
	task := &types.Task{
		Id: uuid.NewV4(),
	}

	testAgent.Operate(task.Id, types.CommandCreate, task)

	taskInManager := testAgent.TaskManager.Get(task.Id)
	assertion.EqualValues(task, taskInManager)
}
