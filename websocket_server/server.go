package websocket_server

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gfile"
)

func Send(client *ghttp.WebSocket, msgType int, data []byte){
	client.WriteMessage(msgType, data)
}

func Start() {
	server := g.Server()
	server.BindHandler("/ws", looper)
	server.SetServerRoot(gfile.MainPkgPath())
	server.SetPort(8199)
	server.Run()
}
