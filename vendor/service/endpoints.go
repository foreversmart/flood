package service

import (
	"github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"
	"log"
)

func MakeOperateEndpoint(as AgentService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(OperateRequest)
		err, resp := as.Operate(req.Id, req.Operate, req.Data)

		log.Println("make startb endpoint")

		if err != nil {
			return OperateResponse{
				Success: false,
				Msg:     err.Error(),
			}, err
		}

		return OperateResponse{
			Success: true,
			Msg:     "ok",
			Data:    resp,
		}, nil
	}

}
