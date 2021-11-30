package router

import (
	"github.com/kanyuanzhi/tialloy/tiface"
	"github.com/kanyuanzhi/tialloy/tinet"
	"tialloy-demo/face"
)

type TcpRouter struct {
	*tinet.BaseRouter
	TrafficHub face.ITrafficHub
}

func NewTcpRouter(trafficHub face.ITrafficHub) tiface.IRouter {
	return &TcpRouter{
		TrafficHub: trafficHub,
	}
}

func (tr *TcpRouter) Handle(request tiface.IRequest) {
	tr.TrafficHub.OnTcpArrive(request)
}