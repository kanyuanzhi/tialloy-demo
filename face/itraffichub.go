package face

import "github.com/kanyuanzhi/tialloy/tiface"

type ITrafficHub interface {
	Start()
	Stop()

	Subscribe(request tiface.IRequest, macs []string)
	UnSubscribe(request tiface.IRequest, mac string)

	OnTcpArrive(request tiface.IRequest)
}