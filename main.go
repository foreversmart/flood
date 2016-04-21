package main

import (
	"code.google.com/p/go.net/context"
	"flood/agent"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/qiniu/log"
	"net/http"
)

func main() {
	ctx := context.Background()
	as := agent.DefaultAgent{}

	startHandler := httptransport.NewServer(
		ctx,
		agent.MakeOperateEndpoint(as),
		agent.DecodeOperateRequest,
		agent.EncodeResponse,
	)

	http.Handle("/operate", startHandler)
	log.Info("test")
	log.Fatal(http.ListenAndServe(":8090", nil))
}
