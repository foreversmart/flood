package executor

import (
	"bytes"
	"io"
	"net/http"
	"types"
)

type _Helper struct{}

var (
	Helper *_Helper
)

func (_ *_Helper) ExecuteContent(content *types.Content) error {
	request, err := Helper.Request(content)
	if err != nil {
		return nil, err
	}

	return http.DefaultClient.Do(request)
}

func (_ *_Helper) Request(content *types.Content) (request *http.Request, err error) {
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
		if len(value) > 0 {
			request.Header.Set(key, value[0])
		}
	}

	return
}
