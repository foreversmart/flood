package executor

import (
	"sync"
	"time"
	"types"
)

type Executor struct {
	ID          string
	Task        *types.Task
	Start       int64
	End         int64
	mux         sync.Mutex // lock for meta info
	concurrence chan bool  // max concurrence execute queue

}

func (executor *Executor) Run() {
	// set max task concurrence for executor
	executor.InitConcurrence()

	// task execute no more than task.repeat times
	for i := 0; i < executor.Task.Repeat; i++ {
		executor.AddTaskToQueue()
		// execute task
		go executor.RunTask()
	}



}

func (executor *Executor) InitConcurrence() {
	executor.concurrence = make(chan bool, executor.Task.Concurrence)
}

func (executor *Executor) AddTaskToQueue() {
	executor.concurrence <- true
}

func (executor *Executor) ReleaseTaskFromQueue() {
	<-executor
}

// how to
func (executor *Executor) RunTask() {
	defer executor.ReleaseTaskFromQueue()

	for _, item := range executor.Task.Items {
		time.Sleep(time.Duration(item.BeforeIdle) * time.Second)

		// over max keep
		if executor.Task.Keep > 0 && int(time.Now().Add(time.Duration(executor.Start)).Second()) > executor.Task.Keep {
			return
		}

		Helper.ExecuteContent(item.Content)

		time.Sleep(time.Duration(item.AfterIdle) * time.Second)
	}
}
