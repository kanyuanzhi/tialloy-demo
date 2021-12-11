package tcp

import (
	"github.com/kanyuanzhi/tialloy/tiface"
	"tialloy-demo/face"
)

type TerminalRunningRouter struct {
	*BaseTcpRouter
}

func NewTerminalRunningRouter(trafficHub face.ITrafficHub) tiface.IRouter {
	return &TerminalRunningRouter{
		NewBaseTcpRouter(trafficHub),
	}
}

func (r *TerminalRunningRouter) Handle(request tiface.IRequest) {
	r.TrafficHub.OnTcpArrive(request)
}
