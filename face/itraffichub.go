package face

import "github.com/kanyuanzhi/tialloy/tiface"

type ITrafficHub interface {
	Start() // 启动
	Stop()  // 停止

	Subscribe(request tiface.IRequest, macs []string)      // websocket request订阅需要的信息，以mac地址为标记
	UnSubscribe(requests chan tiface.IRequest, mac string) // websocket request取消订阅需要的信息，以mac地址为标记

	OnTcpArrive(request tiface.IRequest) // Tcp信息到达时的钩子函数
}
