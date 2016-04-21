package main

import (
	"code.google.com/p/go.net/context"
	"flood/server"
	"kit/log"
	"os"
	"flood/agent"
)

func main() {
	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.NewContext(logger).With("listen", "8090").With("caller", log.DefaultCaller)

	ctx := context.Background()

	var as agent.AgentService
	proxy := server.ProxyingMiddleware("127.0.0.1:8090", ctx, logger)(as)
	proxy.Operate("1231", "start", nil)

}
