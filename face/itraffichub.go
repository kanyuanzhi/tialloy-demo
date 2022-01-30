package face

import "github.com/kanyuanzhi/tialloy/tiface"

type ITrafficHub interface {
	Start() // 启动
	Stop()  // 停止
	Serve()

	OnWebsocketArrive(request tiface.IRequest) //
	OnTcpArrive(request tiface.IRequest)       //

	OnCommandArrive(request tiface.IRequest)
	SetTcpConnList(request tiface.IRequest)

	AddSubscribeList(msgID uint32)
}
