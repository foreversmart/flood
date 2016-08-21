package proxy

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"golang.org/x/net/context"

	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/loadbalancer"
	"github.com/go-kit/kit/loadbalancer/static"
	kitratelimit "github.com/go-kit/kit/ratelimit"
	httptransport "github.com/go-kit/kit/transport/http"
	jujuratelimit "github.com/juju/ratelimit"
	"github.com/sony/gobreaker"

	"service"

	"fmt"
	"io"
	"net/url"
	"strings"
	"time"
)

func ProxyingMiddleware(proxyList string, ctx context.Context, logger log.Logger) service.ServiceMiddleware {
	if proxyList == "" {
		logger.Log("proxy_to", "none")
		return func(next service.AgentService) service.AgentService { return next }
	}
	proxies := split(proxyList)
	logger.Log("proxy_to", fmt.Sprint(proxies))

	return func(next service.AgentService) service.AgentService {
		var (
			qps         = 100 // max to each instance
			publisher   = static.NewPublisher(proxies, factory(ctx, qps), logger)
			lb          = loadbalancer.NewRoundRobin(publisher)
			maxAttempts = 3
			maxTime     = 100 * time.Millisecond
			endpoint    = loadbalancer.Retry(maxAttempts, maxTime, lb)
		)
		return Proxymw{ctx, endpoint, next}
	}
}

// Proxymw implements OperateService, forwarding Uppercase requests to the
// provided endpoint, and serving all other (i.e. Count) requests via the
// embedded service.AgentService.
type Proxymw struct {
	context.Context
	OperateEndpoint      endpoint.Endpoint // ...except Uppercase, which gets served by this endpoint
	service.AgentService                   // Serve most requests via this embedded service...
}

func (mw Proxymw) Operate(id, operate string, data interface{}) (err error) {
	_, err = mw.OperateEndpoint(mw.Context, service.OperateRequest{Id: id, Operate: operate, Data: data})
	return
}

func factory(ctx context.Context, qps int) loadbalancer.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		var e endpoint.Endpoint
		e = makeOperateProxy(ctx, instance)
		e = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(e)
		e = kitratelimit.NewTokenBucketLimiter(jujuratelimit.NewBucketWithRate(float64(qps), int64(qps)))(e)
		return e, nil, nil
	}
}

func makeOperateProxy(ctx context.Context, instance string) endpoint.Endpoint {
	if !strings.HasPrefix(instance, "http") {
		instance = "http://" + instance
	}
	u, err := url.Parse(instance)
	if err != nil {
		panic(err)
	}
	if u.Path == "" {
		u.Path = "/operate"
	}
	return httptransport.NewClient(
		"GET",
		u,
		service.EncodeRequest,
		service.DecodeOperateResponse,
	).Endpoint()
}

func split(s string) []string {
	a := strings.Split(s, ",")
	for i := range a {
		a[i] = strings.TrimSpace(a[i])
	}
	return a
}
