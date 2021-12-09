package tcp

import (
	"github.com/kanyuanzhi/tialloy/tiface"
	"github.com/kanyuanzhi/tialloy/utils"
	"tialloy-demo/face"
)

type TerminalBasicInfoRouter struct {
	*BaseTcpRouter
}

func NewTerminalBasicInfoRouter(trafficHub face.ITrafficHub) tiface.IRouter {
	return &TerminalBasicInfoRouter{
		NewBaseTcpRouter(trafficHub),
	}
}

func (r *TerminalBasicInfoRouter) Handle(request tiface.IRequest) {
	utils.GlobalLog.Info("send basic info to web-service server to register")
	// TODO: send basic info to web-service server to register
}
