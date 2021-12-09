package face

import "github.com/kanyuanzhi/tialloy/tiface"

type ITrafficHub interface {
	Start() // 启动
	Stop()  // 停止

	OnWebsocketArrive(request tiface.IRequest) //
	OnTcpArrive(request tiface.IRequest)       //

	AddSubscribeList(msgID uint32)
}
