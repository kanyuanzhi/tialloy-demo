package tcp

import (
	"github.com/kanyuanzhi/tialloy/tinet"
	"tialloy-demo/face"
)

type BaseTcpRouter struct {
	*tinet.BaseRouter
	TrafficHub face.ITrafficHub
}

func NewBaseTcpRouter(trafficHub face.ITrafficHub) *BaseTcpRouter {
	return &BaseTcpRouter{
		TrafficHub: trafficHub,
	}
}
