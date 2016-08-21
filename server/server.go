package main

import (
	"code.google.com/p/go.net/context"
	httptransport "github.com/go-kit/kit/transport/http"

	"flood/server/agent"
	. "log"
	"net/http"
	"service"
)

func main() {
	server()
}

func server() {
	ctx := context.Background()
	as := agent.DefaultAgent{}

	startHandler := httptransport.NewServer(
		ctx,
		service.MakeOperateEndpoint(as),
		service.DecodeOperateRequest,
		service.EncodeResponse,
	)

	http.Handle("/operate", startHandler)
	Fatal(http.ListenAndServe(":8091", nil))
}
