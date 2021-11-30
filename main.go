package main

import (
	"github.com/kanyuanzhi/tialloy/tinet"
	"time"
)

var TrafficHubIns *TrafficHub

func main() {
	tcpServer := tinet.NewTcpServer()
	tr := NewTcpRouter()
	tcpServer.AddRouter(1, tr)
	go tcpServer.Serve()

	websocketServer := tinet.NewWebsocketServer()
	wr := NewWebsocketRouter()
	websocketServer.AddRouter(1, wr)
	go websocketServer.Serve()

	TrafficHubIns = NewTrafficHub()

	for {
		time.Sleep(time.Minute)
	}
}
