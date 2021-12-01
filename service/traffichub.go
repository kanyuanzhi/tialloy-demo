package service

import (
	"encoding/json"
	"github.com/kanyuanzhi/tialloy/tiface"
	"github.com/kanyuanzhi/tialloy/utils"
	"sync"
	"tialloy-demo/face"
	"tialloy-demo/model"
)

type TrafficHub struct {
	WebsocketServer tiface.IServer
	TcpServer       tiface.IServer

	SubscribeList     map[string]map[tiface.IRequest]bool // 针对单个mac地址所标记信息的订阅列表
	SubscribeListLock *sync.RWMutex

	TcpArrivalChan chan []byte
}

func NewTrafficHub(websocketServer tiface.IServer, tcpServer tiface.IServer) face.ITrafficHub {
	trafficHub := &TrafficHub{
		WebsocketServer:   websocketServer,
		TcpServer:         tcpServer,
		SubscribeList:     make(map[string]map[tiface.IRequest]bool),
		SubscribeListLock: new(sync.RWMutex),
		TcpArrivalChan:    make(chan []byte, 20),
	}
	return trafficHub
}

func (th *TrafficHub) Start() {
	for {
		select {
		case tcpMessageData := <-th.TcpArrivalChan:
			go func(tcpMessageData []byte) {
				var status = model.Status{}
				_ = json.Unmarshal(tcpMessageData, &status)
				unsubscribeRequests := make(chan tiface.IRequest, 100)
				if _, ok := th.SubscribeList[status.Mac]; ok {
					wg := sync.WaitGroup{}
					wg.Add(len(th.SubscribeList[status.Mac]))
					for websocketRequest, _ := range th.SubscribeList[status.Mac] {
						go func(websocketRequest tiface.IRequest) {
							err := websocketRequest.GetConnection().SendBuffMsg(websocketRequest.GetMsgID(), tcpMessageData)
							if err != nil {
								utils.GlobalLog.Error(err)
								unsubscribeRequests <- websocketRequest
							}
							wg.Done()
						}(websocketRequest)
					}
					wg.Wait()
					th.UnSubscribe(unsubscribeRequests, status.Mac)
					utils.GlobalLog.Tracef("there is(are) %d websocket request(s) for %s", len(th.SubscribeList[status.Mac]), status.Mac)
				} else {
					utils.GlobalLog.Tracef("there are no websocket requests for %s", status.Mac)
					return
				}
			}(tcpMessageData)
		}
	}
}

func (th *TrafficHub) Stop() {
	panic("implement me")
}

func (th *TrafficHub) SubscribeOne(mac string, request tiface.IRequest) {
	th.SubscribeListLock.Lock()
	defer th.SubscribeListLock.Unlock()
	if _, ok := th.SubscribeList[mac]; !ok {
		th.SubscribeList[mac] = make(map[tiface.IRequest]bool)
		th.SubscribeList[mac][request] = true
	} else {
		th.SubscribeList[mac][request] = true
	}
}

func (th *TrafficHub) Subscribe(request tiface.IRequest, macs []string) {
	for _, mac := range macs {
		th.SubscribeOne(mac, request)
	}
}

func (th *TrafficHub) UnSubscribe(requests chan tiface.IRequest, mac string) {
	th.SubscribeListLock.Lock()
	defer th.SubscribeListLock.Unlock()
	close(requests)
	for request := range requests {
		delete(th.SubscribeList[mac], request)
		utils.GlobalLog.Warnf("websocket request via connID=%d is removed from SubscribeList of %s", request.GetConnection().GetConnID(), mac)
	}

}

func (th *TrafficHub) OnTcpArrive(request tiface.IRequest) {
	th.TcpArrivalChan <- request.GetData()
}
