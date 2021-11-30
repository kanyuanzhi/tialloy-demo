package server

import (
	"encoding/json"
	"github.com/kanyuanzhi/tialloy/tiface"
	"github.com/kanyuanzhi/tialloy/utils"
	"tialloy-demo/face"
	"tialloy-demo/model"
)

type TrafficHub struct {
	WebsocketServer tiface.IServer
	TcpServer       tiface.IServer

	SubscribeList        map[string][]tiface.IRequest

	TcpArrivalChan       chan tiface.IRequest
}

func NewTrafficHub(websocketServer tiface.IServer, tcpServer tiface.IServer) face.ITrafficHub {
	trafficHub := &TrafficHub{
		WebsocketServer: websocketServer,
		TcpServer:       tcpServer,
		SubscribeList:   make(map[string][]tiface.IRequest),
		TcpArrivalChan:  make(chan tiface.IRequest, 20),
	}
	return trafficHub
}

func (th *TrafficHub) Start() {
	for {
		select {
		case tcpRequest := <-th.TcpArrivalChan:
			var status = model.Status{}
			data := tcpRequest.GetData()
			_ = json.Unmarshal(data, &status)
			if _, ok := th.SubscribeList[status.Mac]; ok {
				utils.GlobalLog.Tracef("%s is existed", status.Mac)
				for _, websocketRequest := range th.SubscribeList[status.Mac] {
					go func(websocketRequest tiface.IRequest) {
						// TODO: websocket connection关闭时从SubscribeList删除对应request
						err := websocketRequest.GetConnection().SendBuffMsg(websocketRequest.GetMsgID(), data)
						if err != nil {
							utils.GlobalLog.Error(err)
						}
					}(websocketRequest)
				}
			} else {
				utils.GlobalLog.Tracef("%s is not existed", status.Mac)
				continue
			}
		}
	}
}

func (th *TrafficHub) Stop() {
	panic("implement me")
}

func (th *TrafficHub) SubscribeOne(mac string, request tiface.IRequest) {
	if _, ok := th.SubscribeList[mac]; !ok {
		th.SubscribeList[mac] = []tiface.IRequest{request}
	} else {
		th.SubscribeList[mac] = append(th.SubscribeList[mac], request)
	}
}

func (th *TrafficHub) Subscribe(request tiface.IRequest, macs []string) {
	for _, mac := range macs {
		th.SubscribeOne(mac, request)
	}
}

func (th *TrafficHub) UnSubscribe(request tiface.IRequest, mac string) {
	panic("implement me!")
}

func (th *TrafficHub) OnTcpArrive(request tiface.IRequest) {
	utils.GlobalLog.Trace(request)
	th.TcpArrivalChan <- request
}