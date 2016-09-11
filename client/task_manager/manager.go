package task_manager

import (
	"sync"
	"types"
)

type TaskManager struct {
	Tasks map[string]*types.Task
	mutex sync.RWMutex
}

func NewTaskManager() *TaskManager {
	return &TaskManager{
		Tasks: make(map[string]*types.Task),
	}
}

func (m *TaskManager) Add(task *types.Task) {
	m.mutex.Lock()
	m.Tasks[task.Id] = task
	m.mutex.Unlock()
}

func (m *TaskManager) Get(id string) *types.Task {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.Tasks[id]
}

func (m *TaskManager) Delete(id string) {
	m.mutex.Lock()
	delete(m.Tasks, id)
	m.mutex.Unlock()
}
