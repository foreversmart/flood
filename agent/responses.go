package agent

import (
	"encoding/json"
	"log"
	"net/http"
)

type OperateResponse struct {
	Success bool        `json:"success"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func EncodeResponse(w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func DecodeOperateResponse(r *http.Response) (interface{}, error) {
	var response OperateResponse
	log.Println(r)
	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}
