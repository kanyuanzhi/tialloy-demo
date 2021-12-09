package websocket

import (
	"github.com/kanyuanzhi/tialloy/tiface"
	"github.com/kanyuanzhi/tialloy/tinet"
	"tialloy-demo/face"
)

type BaseRouter struct {
	*tinet.BaseRouter
	TrafficHub face.ITrafficHub
}

func NewBaseRouter(trafficHub face.ITrafficHub) tiface.IRouter {
	return &BaseRouter{
		TrafficHub: trafficHub,
	}
}

func (wr *BaseRouter) Handle(request tiface.IRequest) {
	wr.TrafficHub.OnWebsocketArrive(request)
}
