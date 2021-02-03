package websocket_server

import (
	"fmt"
	"github.com/gogf/gf/net/ghttp"
)

func OnConnected(client *ghttp.WebSocket){
	SetRecvDeadline(client, 2)

	fmt.Println("OnConnected")
}

func OnData(client *ghttp.WebSocket, msgType int, data []byte){
	fmt.Println("OnData ", msgType, data)
}

func OnServerKicked(client *ghttp.WebSocket){
	fmt.Println("OnServerKicked")
}

func OnClientDisconected(client *ghttp.WebSocket){
	fmt.Println("OnClientDisconected")
}

func OnTimeOut(client *ghttp.WebSocket){
	fmt.Println("OnTimeOut")
}

