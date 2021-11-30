package router

import (
	"encoding/json"
	"github.com/kanyuanzhi/tialloy/tiface"
	"github.com/kanyuanzhi/tialloy/tinet"
	"github.com/kanyuanzhi/tialloy/utils"
	"tialloy-demo/face"
	"tialloy-demo/model"
)

type WebsocketRouter struct {
	*tinet.BaseRouter
	TrafficHub face.ITrafficHub
}

func NewWebsocketRouter(trafficHub face.ITrafficHub) tiface.IRouter {
	return &WebsocketRouter{
		TrafficHub: trafficHub,
	}
}

func (wr *WebsocketRouter) Handle(request tiface.IRequest) {
	utils.GlobalLog.Infof("msgID=%d, data=%s", request.GetMsgID(), request.GetData())
	var wrd = model.WebsocketRequestModel{}
	_ = json.Unmarshal(request.GetData(), &wrd)
	wr.TrafficHub.Subscribe(request, wrd.Data)
}
