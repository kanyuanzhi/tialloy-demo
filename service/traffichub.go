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

	SubscribeList     map[uint32]map[string]map[tiface.IRequest]bool // {MsgID:{keys:{}}}
	SubscribeListLock sync.RWMutex

	WebsocketArrivalChan chan tiface.IRequest
	TcpArrivalChan       chan tiface.IRequest
}

func NewTrafficHub(websocketServer tiface.IServer, tcpServer tiface.IServer) face.ITrafficHub {
	trafficHub := &TrafficHub{
		WebsocketServer:      websocketServer,
		TcpServer:            tcpServer,
		SubscribeList:        make(map[uint32]map[string]map[tiface.IRequest]bool),
		WebsocketArrivalChan: make(chan tiface.IRequest, 100),
		TcpArrivalChan:       make(chan tiface.IRequest, 100),
	}
	return trafficHub
}

func (th *TrafficHub) Start() {
	for {
		select {
		case tcpReq := <-th.TcpArrivalChan:
			if th.MsgIDCheck(tcpReq) {
				go th.Distribute(tcpReq)
				go th.Save(tcpReq)
			}
		case wsReq := <-th.WebsocketArrivalChan:
			if th.MsgIDCheck(wsReq) {
				go th.Subscribe(wsReq)
			}
		}
	}
}

func (th *TrafficHub) Stop() {
	panic("implement me")
}

func (th *TrafficHub) OnWebsocketArrive(request tiface.IRequest) {
	th.WebsocketArrivalChan <- request
}

func (th *TrafficHub) OnTcpArrive(request tiface.IRequest) {
	th.TcpArrivalChan <- request
}

func (th *TrafficHub) AddSubscribeList(msgID uint32) {
	th.SubscribeList[msgID] = make(map[string]map[tiface.IRequest]bool)
}

func (th *TrafficHub) Distribute(request tiface.IRequest) {
	var tcpReqModel = model.TcpRequest{}
	if err := json.Unmarshal(request.GetData(), &tcpReqModel); err != nil {
		utils.GlobalLog.Error(err)
		return
	}
	unsubscribeWsReq := make(chan tiface.IRequest, 100) //需要取消订阅的已断开连接的websocket集合
	if wsReqs, ok := th.SubscribeList[request.GetMsgID()][tcpReqModel.Key]; ok {
		wg := sync.WaitGroup{}
		wg.Add(len(wsReqs))
		for wsReq, _ := range wsReqs {
			go func(wsReq tiface.IRequest) {
				err := wsReq.GetConnection().SendBuffMsg(request.GetMsgID(), request.GetData())
				if err != nil {
					utils.GlobalLog.Error(err)
					unsubscribeWsReq <- wsReq
				}
				wg.Done()
			}(wsReq)
		}
		wg.Wait()
		th.UnSubscribe(unsubscribeWsReq, request.GetMsgID(), tcpReqModel.Key)
		utils.GlobalLog.Tracef("there is(are) %d websocket request(s) for msgID=%d, key=%s", len(wsReqs), request.GetMsgID(), tcpReqModel.Key)
	} else {
		utils.GlobalLog.Tracef("there are no websocket requests for msgID=%d, key=%s", request.GetMsgID(), tcpReqModel.Key)
		return
	}
}

func (th *TrafficHub) Save(request tiface.IRequest) {
	utils.GlobalLog.Tracef("save msgID=%d", request.GetMsgID())
	// TODO save data
}

func (th *TrafficHub) SubscribeOne(request tiface.IRequest, key string) {
	th.SubscribeListLock.Lock()
	defer th.SubscribeListLock.Unlock()
	if _, ok := th.SubscribeList[request.GetMsgID()][key]; ok {
		th.SubscribeList[request.GetMsgID()][key][request] = true
	} else {
		th.SubscribeList[request.GetMsgID()][key] = make(map[tiface.IRequest]bool)
		th.SubscribeList[request.GetMsgID()][key][request] = true
	}
}

func (th *TrafficHub) Subscribe(request tiface.IRequest) {
	var wrd = model.WebsocketRequest{}
	err := json.Unmarshal(request.GetData(), &wrd)
	if err != nil {
		utils.GlobalLog.Warn(err)
	}
	for _, key := range wrd.Data {
		th.SubscribeOne(request, key)
	}
}

func (th *TrafficHub) UnSubscribe(requests chan tiface.IRequest, msgID uint32, key string) {
	th.SubscribeListLock.Lock()
	defer th.SubscribeListLock.Unlock()
	close(requests)
	for request := range requests {
		delete(th.SubscribeList[msgID][key], request)
		utils.GlobalLog.Warnf("websocket request via connID=%d is removed from SubscribeList of msgID=%d, key=%s", request.GetConnection().GetConnID(), msgID, key)
	}
}

func (th *TrafficHub) MsgIDCheck(request tiface.IRequest) bool {
	_, ok := th.SubscribeList[request.GetMsgID()]
	utils.GlobalLog.Warnf("%s request msgID=%d is not added to the SubscribeList, refuse serving", request.GetConnection().GetServer().GetServerType(), request.GetMsgID())
	return ok
}
