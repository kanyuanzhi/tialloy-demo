package main

import (
	"github.com/kanyuanzhi/tialloy/tinet"
	"tialloy-demo/router"
	"tialloy-demo/service"
	"time"
)

func main() {
	tcpServer := tinet.NewTcpServer()
	websocketServer := tinet.NewWebsocketServer()
	trafficHub := service.NewTrafficHub(websocketServer, tcpServer)

	tcpRouter := router.NewTcpRouter(trafficHub)
	tcpServer.AddRouter(1, tcpRouter)

	websocketRouter := router.NewWebsocketRouter(trafficHub)
	websocketServer.AddRouter(1, websocketRouter)

	go websocketServer.Serve()
	go tcpServer.Serve()
	go trafficHub.Start()

	for {
		time.Sleep(time.Minute)
	}
}
