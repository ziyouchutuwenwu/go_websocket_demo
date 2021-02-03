package websocket_server

import (
	"github.com/gogf/gf/net/ghttp"
	"time"
)

func SetRecvDeadline(client *ghttp.WebSocket, second time.Duration){
	client.SetReadDeadline(time.Now().Add(time.Second * second))
}

func SetSendDeadline(client *ghttp.WebSocket, second time.Duration){
	client.SetWriteDeadline(time.Now().Add(time.Second * second))
}