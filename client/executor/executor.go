package executor

import (
	"bytes"
	"heka/build/heka/src/code.google.com/p/gogoprotobuf/io"
	"net/http"
	"types"
)

type _Executor struct{}

var (
	Executor *_Executor
)

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
