package types

import "net/http"

type Task struct {
	BeforeIdle  int64   `json:"before_idle"`
	AfterIdle   int64   `json:"after_idle"`
	Items       []*Item `json:"items"`
	Repeat      int     `json:"repeat"`
	Keep        int     `json:"keep"`
	Concurrence int     `json:"concurrence"`
}

// execution item describe the request content
type Item struct {
	BeforeIdle int64   `json:"before_idle"`
	AfterIdle  int64   `json:"after_idle"`
	Content    Content `json:"content"`
}

// execution request content
type Content struct {
	Url      string      `json:"url"`
	Header   http.Header `json:"header"`
	Method   string      `json:"method"`
	Body     []byte      `json:"body"`
	DataType string      `json:"data_type"`
}
