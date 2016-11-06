package agent

import (
	"kit/log"
	"os"
	"proxy"
	"service"

	"code.google.com/p/go.net/context"
)

var (
	Server *service.AgentService
)

func init() {
	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.NewContext(logger).With("listen", "8091").With("caller", log.DefaultCaller)

	ctx := context.Background()

	var as service.AgentService
	Server = proxy.ProxyingMiddleware("127.0.0.1:8091", ctx, logger)(as)
}
