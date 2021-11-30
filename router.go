package main

import (
	"encoding/json"
	"github.com/kanyuanzhi/tialloy/tiface"
	"github.com/kanyuanzhi/tialloy/tinet"
	"github.com/kanyuanzhi/tialloy/utils"
)

type WebsocketRequest struct {
	MsgID uint32   `json:"msg_id,omitempty"`
	Data  []string `json:"data,omitempty"`
}

type Status struct {
	Mac    string `json:"mac,omitempty"`
	Number uint32 `json:"number,omitempty"`
}

type WebsocketRouter struct {
	*tinet.BaseRouter
}

func NewWebsocketRouter() tiface.IRouter {
	return &WebsocketRouter{}
}

func (wr *WebsocketRouter) Handle(request tiface.IRequest) {
	utils.GlobalLog.Infof("msgID=%d, data=%s", request.GetMsgID(), request.GetData())
	var wsReq = WebsocketRequest{}
	_ = json.Unmarshal(request.GetData(), &wsReq)
	//request.GetConnection().SendBuffMsg(request.GetMsgID(), []byte("hello"))

	trafficHub.Subscribe(wsReq.Data, request.GetConnection().(*tinet.WebsocketConnection))
}

type TcpRouter struct {
	*tinet.BaseRouter
}

func NewTcpRouter() tiface.IRouter {
	return &TcpRouter{}
}

func (tr *TcpRouter) Handle(request tiface.IRequest) {
	var status = Status{}
	data := request.GetData()
	_ = json.Unmarshal(data, &status)
	utils.GlobalLog.Info(status)
}
