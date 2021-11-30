package main

import (
	"github.com/kanyuanzhi/tialloy/tinet"
	"sync"
)

type TrafficHub struct {
	//TrafficChanIn  map[string]chan interface{}
	//TrafficChanOut map[string]chan interface{}

	SubscribeList map[string][]*tinet.WebsocketConnection
	SubscribeLock *sync.Mutex
}

func NewTrafficHub() *TrafficHub {
	return &TrafficHub{
		SubscribeList: make(map[string][]*tinet.WebsocketConnection),
	}
}

func (th *TrafficHub) SubscribeOne(mac string, wsConnection *tinet.WebsocketConnection) {
	if _, ok := th.SubscribeList[mac]; !ok {
		th.SubscribeList[mac] = []*tinet.WebsocketConnection{wsConnection}
	} else {
		th.SubscribeList[mac] = append(th.SubscribeList[mac], wsConnection)
	}
}

func (th *TrafficHub) Subscribe(macs []string, wsConnection *tinet.WebsocketConnection) {
	for _, mac := range macs {
		th.SubscribeLock.Lock()
		th.SubscribeOne(mac, wsConnection)
		th.SubscribeLock.Unlock()
	}
}
