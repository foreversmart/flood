package agent

import (
	"kit/log"
	"os"
	"proxy"
	"service"

	"golang.org/x/net/context"
	"types"
)

var (
	Client service.AgentService
)

func init() {
	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.NewContext(logger).With("listen", "8090").With("caller", log.DefaultCaller)

	ctx := context.Background()

	//var as service.AgentService
	var as service.AgentService
	Client = proxy.ProxyingMiddleware("127.0.0.1:8090", ctx, logger)(as)
}

func Start(task *types.Task) error {
	err, _ := Client.Operate(task.Id, types.CommandStart, task)
	return err
}
