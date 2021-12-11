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

	tcpTerminalBasicRouter := tcp.NewTerminalBasicRouter(trafficHub)
	tcpServer.AddRouter(101, tcpTerminalBasicRouter)
	trafficHub.AddSubscribeList(101)

	tcpTerminalRunningRouter := tcp.NewTerminalRunningRouter(trafficHub)
	tcpServer.AddRouter(102, tcpTerminalRunningRouter)
	trafficHub.AddSubscribeList(102)

	websocketBaseRouter := websocket.NewBaseRouter(trafficHub)
	websocketServer.AddRouter(102, websocketBaseRouter)

	go websocketServer.Serve()
	go tcpServer.Serve()
	go trafficHub.Serve()

	select {}
}
