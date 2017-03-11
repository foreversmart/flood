package types

import "net/http"

type Task struct {
	Id          string  `json:"id"`
	BeforeIdle  int64   `json:"before_idle"`
	AfterIdle   int64   `json:"after_idle"`
	Items       []*Item `json:"items"`
	Repeat      int     `json:"repeat"`      // task max repeat times
	Keep        int     `json:"keep"`        // task max keep alive time
	Concurrence int     `json:"concurrence"` // task max concurrence
	RateLimit   int     `json:"rate_limit"`

	TaskState TaskState `json:"task_state"`
}

type TaskState int

const (
	TaskIdle TaskState = iota
	TaskPendding
	TaskRunning
	TaskStopped
)

// execution item describe the request content
type Item struct {
	BeforeIdle int64    `json:"before_idle"`
	AfterIdle  int64    `json:"after_idle"`
	Content    *Content `json:"content"`
}

// execution request content
type Content struct {
	Url      string      `json:"url"`
	Header   http.Header `json:"header"`
	Method   string      `json:"method"`
	Body     []byte      `json:"body"`
	DataType string      `json:"data_type"`
}
