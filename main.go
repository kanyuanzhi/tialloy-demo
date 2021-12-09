package main

import (
	"github.com/kanyuanzhi/tialloy/tinet"
	"tialloy-demo/router/tcp"
	"tialloy-demo/router/websocket"
	"tialloy-demo/service"
)

func main() {
	tcpServer := tinet.NewTcpServer()
	websocketServer := tinet.NewWebsocketServer()
	trafficHub := service.NewTrafficHub(websocketServer, tcpServer)

	tcpTerminalBasicInfoRouter := tcp.NewTerminalBasicInfoRouter(trafficHub)
	tcpServer.AddRouter(101, tcpTerminalBasicInfoRouter)
	trafficHub.AddSubscribeList(101)

	tcpTerminalRunningInfoRouter := tcp.NewTerminalRunningInfoRouter(trafficHub)
	tcpServer.AddRouter(102, tcpTerminalRunningInfoRouter)
	trafficHub.AddSubscribeList(102)

	websocketBaseRouter := websocket.NewBaseRouter(trafficHub)
	websocketServer.AddRouter(102, websocketBaseRouter)

	go websocketServer.Serve()
	go tcpServer.Serve()
	go trafficHub.Start()

	select {}
}
