package websocket

import (
	"github.com/kanyuanzhi/tialloy/tiface"
	"tialloy-demo/face"
)

type CommandRouter struct {
	*BaseWebsocketRouter
}

func NewCommandRouter(trafficHub face.ITrafficHub) tiface.IRouter {
	return &CommandRouter{
		NewBaseWebsocketRouter(trafficHub),
	}
}

func (cr *CommandRouter) Handle(request tiface.IRequest) {
	cr.TrafficHub.OnCommandArrive(request)
}
