package main

import (
	"github.com/kanyuanzhi/tialloy/tiface"
	"github.com/kanyuanzhi/tialloy/tinet"
	"github.com/kanyuanzhi/tialloy/utils"
)

type WebsocketRouter struct {
	*tinet.BaseRouter
}

func NewWebsocketRouter() tiface.IRouter {
	return &WebsocketRouter{}
}

func (wr *WebsocketRouter) Handle(request tiface.IRequest) {
	utils.GlobalLog.Infof("msgID=%d, data=%s", request.GetMsgID(), request.GetData())
	request.GetConnection().SendBuffMsg(12, []byte("hello"))
}

type TcpRouter struct {
	*tinet.BaseRouter
}

func NewTcpRouter() tiface.IRouter {
	return &TcpRouter{}
}

func (tr *TcpRouter) Handle(request tiface.IRequest) {

}
