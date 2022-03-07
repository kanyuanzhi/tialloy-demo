package service

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"sync"
	"tialloy-demo/face"
	"tialloy-demo/model"

	"github.com/kanyuanzhi/tialloy/tiface"
	"github.com/kanyuanzhi/tialloy/tilog"
)

type TrafficHub struct {
	WebsocketServer tiface.IServer
	TcpServer       tiface.IServer

	SubscribeList     map[uint32]map[string]map[tiface.IRequest]bool // {MsgID:{keys:{}}}
	SubscribeListLock sync.RWMutex

	WebsocketArrivalChan chan tiface.IRequest
	TcpArrivalChan       chan tiface.IRequest

	TcpConnList        map[string]tiface.IConnection
	CommandArrivalChan chan tiface.IRequest
}

func NewTrafficHub(websocketServer tiface.IServer, tcpServer tiface.IServer) face.ITrafficHub {
	trafficHub := &TrafficHub{
		WebsocketServer:      websocketServer,
		TcpServer:            tcpServer,
		SubscribeList:        make(map[uint32]map[string]map[tiface.IRequest]bool),
		WebsocketArrivalChan: make(chan tiface.IRequest, 100),
		TcpArrivalChan:       make(chan tiface.IRequest, 100),

		TcpConnList:        make(map[string]tiface.IConnection),
		CommandArrivalChan: make(chan tiface.IRequest, 100),
	}
	return trafficHub
}

func (th *TrafficHub) Start() {
	go func() {
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
			case commandReq := <-th.CommandArrivalChan:
				go th.SendCommands(commandReq)
			}
		}
	}()
}

func (th *TrafficHub) Serve() {
	th.Start()
	select {}
}

func (th *TrafficHub) Stop() {
	panic("implement me")
}

func (th *TrafficHub) OnWebsocketArrive(request tiface.IRequest) {
	th.WebsocketArrivalChan <- request
}

func (th *TrafficHub) OnCommandArrive(request tiface.IRequest) {
	th.CommandArrivalChan <- request
}

func (th *TrafficHub) SetTcpConnList(request tiface.IRequest) {
	var tcpReqModel = model.TcpRequest{}
	if err := json.Unmarshal(request.GetData(), &tcpReqModel); err != nil {
		tilog.Log.Error(err)
		return
	}
	th.TcpConnList[tcpReqModel.Key] = request.GetConnection()
	tilog.Log.Infoln(th.TcpConnList)
}

func (th *TrafficHub) OnTcpArrive(request tiface.IRequest) {
	tilog.Log.Tracef("OnTcpArrive")
	th.TcpArrivalChan <- request
}

func (th *TrafficHub) AddSubscribeList(msgID uint32) {
	th.SubscribeList[msgID] = make(map[string]map[tiface.IRequest]bool)
}

func (th *TrafficHub) Distribute(request tiface.IRequest) {
	var keyInSubscribeList string
	msgID := request.GetMsgID()
	if msgID == 121 {
		cameraKeyBytes := request.GetData()[:4]
		dataBuf := bytes.NewBuffer(cameraKeyBytes)
		var cameraKey uint32
		err := binary.Read(dataBuf, binary.BigEndian, &cameraKey)
		if err != nil {
			tilog.Log.Error(err)
		}
		keyInSubscribeList = fmt.Sprintf("%d", cameraKey)
	} else {
		var tcpReqModel = model.TcpRequest{}
		if err := json.Unmarshal(request.GetData(), &tcpReqModel); err != nil {
			tilog.Log.Error(err)
			return
		}
		keyInSubscribeList = tcpReqModel.Key
	}
	unsubscribeWsReq := make(chan tiface.IRequest, 100) //需要取消订阅的已断开连接的websocket集合
	if wsReqs, ok := th.SubscribeList[request.GetMsgID()][keyInSubscribeList]; ok {
		wg := sync.WaitGroup{}
		wg.Add(len(wsReqs))
		for wsReq, _ := range wsReqs {
			go func(wsReq tiface.IRequest) {
				err := wsReq.GetConnection().SendBuffMsg(request.GetMsgID(), request.GetData())
				if err != nil {
					tilog.Log.Error(err)
					unsubscribeWsReq <- wsReq
				}
				wg.Done()
			}(wsReq)
		}
		wg.Wait()
		th.UnSubscribe(unsubscribeWsReq, request.GetMsgID(), keyInSubscribeList)
		tilog.Log.Tracef("there is(are) %d websocket request(s) for msgID=%d, key=%s", len(wsReqs), request.GetMsgID(), keyInSubscribeList)
	} else {
		tilog.Log.Tracef("there are no websocket requests for msgID=%d, key=%s", request.GetMsgID(), keyInSubscribeList)
		return
	}
}

func (th *TrafficHub) Save(request tiface.IRequest) {
	tilog.Log.Tracef("save msgID=%d", request.GetMsgID())
	// TODO save running data
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
		tilog.Log.Warn(err)
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
		tilog.Log.Warnf("websocket request via connID=%d is removed from SubscribeList of msgID=%d, key=%s", request.GetConnection().GetConnID(), msgID, key)
	}
}

func (th *TrafficHub) SendOneCommand(request tiface.IRequest, key string, command string) error {
	tcpCommandRequest := &model.TcpCommandRequest{Command: command}
	data, _ := json.Marshal(tcpCommandRequest)
	if tcpConn, ok := th.TcpConnList[key]; ok {
		err := tcpConn.SendMsg(request.GetMsgID(), data)
		return err
	}
	return errors.New("offline error")
}

func (th *TrafficHub) SendCommands(request tiface.IRequest) {
	var wcr = model.WebsocketCommandRequest{}
	err := json.Unmarshal(request.GetData(), &wcr)
	if err != nil {
		tilog.Log.Warn(err)
	}
	unServedKeys := []string{}
	tilog.Log.Infoln(wcr.Command)
	for _, key := range wcr.Data {
		if err = th.SendOneCommand(request, key, wcr.Command); err != nil {
			unServedKeys = append(unServedKeys, key)
		}
	}
	wsRes := &model.WebsocketResponse{
		MsgID: request.GetMsgID(),
		Data:  unServedKeys,
	}
	data, _ := json.Marshal(wsRes)
	err = request.GetConnection().SendMsg(request.GetMsgID(), data)
	if err != nil {
		tilog.Log.Warn(err)
	}
}

func (th *TrafficHub) MsgIDCheck(request tiface.IRequest) bool {
	var ok bool
	if _, ok = th.SubscribeList[request.GetMsgID()]; !ok {
		tilog.Log.Warnf("%s request msgID=%d is not added to the SubscribeList, refuse serving", request.GetConnection().GetServer().GetServerType(), request.GetMsgID())
	}
	return ok
}
