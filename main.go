package main

import (
	"tialloy-demo/router/tcp"
	"tialloy-demo/router/websocket"
	"tialloy-demo/service"

	"github.com/kanyuanzhi/tialloy/tinet"
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

	tcpImageRouter := tcp.NewBaseTcpRouter(trafficHub)
	tcpServer.AddRouter(120, tcpImageRouter)
	trafficHub.AddSubscribeList(120)

	// 订阅终端运行状态
	subscribeTerminalRouter := websocket.NewBaseWebsocketRouter(trafficHub)
	websocketServer.AddRouter(102, subscribeTerminalRouter)
	// 订阅命令执行情况
	subscribeCommandRouter := websocket.NewBaseWebsocketRouter(trafficHub)
	websocketServer.AddRouter(110, subscribeCommandRouter)
	// 发送命令
	sendCommandRouter := websocket.NewCommandRouter(trafficHub)
	websocketServer.AddRouter(111, sendCommandRouter)

	// 订阅监控画面
	subscribeImageRouter := websocket.NewBaseWebsocketRouter(trafficHub)
	websocketServer.AddRouter(120, subscribeImageRouter)

	go websocketServer.Serve()
	go tcpServer.Serve()
	go trafficHub.Serve()

	select {}
}
