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

	tcpCommandExexResultRouter := tcp.NewBaseTcpRouter(trafficHub)
	tcpServer.AddRouter(110, tcpCommandExexResultRouter)
	trafficHub.AddSubscribeList(110)

	// 订阅终端运行状态
	subscribeTerminalRouter := websocket.NewBaseWebsocketRouter(trafficHub)
	websocketServer.AddRouter(102, subscribeTerminalRouter)
	// 订阅命令执行情况
	subscribeCommandRouter := websocket.NewBaseWebsocketRouter(trafficHub)
	websocketServer.AddRouter(110, subscribeCommandRouter)
	// 发送命令
	sendCommandRouter := websocket.NewCommandRouter(trafficHub)
	websocketServer.AddRouter(111, sendCommandRouter)

	go websocketServer.Serve()
	go tcpServer.Serve()
	go trafficHub.Serve()

	select {}
}
