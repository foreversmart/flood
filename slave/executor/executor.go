package executor

import (
	"bytes"
	"heka/build/heka/src/code.google.com/p/gogoprotobuf/io"
	"net/http"
	"time"
	"types"
)

type _Executor struct{}

var (
	Executor *_Executor
)

func (_ *_Executor) Run(task *types.Task) {
	for i := 0; i < task.Concurrence; i++ {
		go Executor.run(task)
	}

}

func (_ *_Executor) run(task *types.Task) {
	start := time.Now()
	for _, item := range task.Items {
		time.Sleep(item.BeforeIdle)

		// over max keep
		if task.Keep > 0 && int(time.Now().Sub(start).Seconds()) > task.Keep {
			return
		}

		Executor.ExecuteContent(item.Content)

		time.Sleep(item.AfterIdle)
	}
}

func (_ *_Executor) ExecuteContent(content *types.Content) (*http.Response, error) {
	request, err := Executor.Request(content)
	if err != nil {
		return nil, err
	}

	return http.DefaultClient.Do(request)
}

func (_ *_Executor) Request(content *types.Content) (request *http.Request, err error) {
	var reader io.Reader
	if len(content.Body) > 0 {
		reader = bytes.NewBuffer(content.Body)
	}

	request, err = http.NewRequest(content.Method, content.Url, reader)
	if err != nil {
		return
	}

	// set headers
	for key, value := range content.Header {
		request.Header.Set(key, value)
	}

	return
}
