package service

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"types"
)

type OperateRequest struct {
	Id      string            `json:"id"`
	Operate types.CommandType `json:"operate"`
	Data    interface{}       `json:"data"`
}

func DecodeOperateRequest(r *http.Request) (interface{}, error) {
	log.Println("decode start request")
	var request OperateRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return request, nil
}

func EncodeRequest(r *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}
