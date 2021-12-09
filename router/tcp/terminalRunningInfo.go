package tcp

import (
	"github.com/kanyuanzhi/tialloy/tiface"
	"tialloy-demo/face"
)

type TerminalRunningInfoRouter struct {
	*BaseTcpRouter
}

func NewTerminalRunningInfoRouter(trafficHub face.ITrafficHub) tiface.IRouter {
	return &TerminalRunningInfoRouter{
		NewBaseTcpRouter(trafficHub),
	}
}

func (r *TerminalRunningInfoRouter) Handle(request tiface.IRequest) {
	r.TrafficHub.OnTcpArrive(request)
}
