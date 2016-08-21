package main

import (
	"code.google.com/p/go.net/context"
	httptransport "github.com/go-kit/kit/transport/http"

	"log"
	"net/http"

	"flood/client/agent"
	"service"
)

func main() {
	go clientServer()
}

func clientServer() {
	ctx := context.Background()
	as := agent.DefaultAgent{}

	startHandler := httptransport.NewServer(
		ctx,
		service.MakeOperateEndpoint(as),
		service.DecodeOperateRequest,
		service.EncodeResponse,
	)

	http.Handle("/operate", startHandler)
	log.Fatal(http.ListenAndServe(":8090", nil))
}
