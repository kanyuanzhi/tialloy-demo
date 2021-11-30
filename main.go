package main

import (
	"github.com/kanyuanzhi/tialloy/tinet"
	"time"
)

func main()  {
	tcpServer := tinet.NewTcpServer()
	go tcpServer.Serve()

	websocketServer := tinet.NewWebsocketServer()
	wr := NewWebsocketRouter()
	websocketServer.AddRouter(1, wr)
	go websocketServer.Serve()

	for {
		time.Sleep(time.Minute)
	}
}
