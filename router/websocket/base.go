package websocket

import (
	"github.com/kanyuanzhi/tialloy/tiface"
	"github.com/kanyuanzhi/tialloy/tinet"
	"tialloy-demo/face"
)

type BaseWebsocketRouter struct {
	*tinet.BaseRouter
	TrafficHub face.ITrafficHub
}

func NewBaseWebsocketRouter(trafficHub face.ITrafficHub) *BaseWebsocketRouter {
	return &BaseWebsocketRouter{
		TrafficHub: trafficHub,
	}
}

func (bwr *BaseWebsocketRouter) Handle(request tiface.IRequest) {
	bwr.TrafficHub.OnWebsocketArrive(request)
}
