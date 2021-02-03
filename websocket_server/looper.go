package websocket_server

import (
	"github.com/gogf/gf/net/ghttp"
	"strings"
)

func looper(request *ghttp.Request){
	client, _ := request.WebSocket()
	OnConnected(client)

	for {
		msgType, data, err := client.ReadMessage()
		if err == nil{
			OnData(client, msgType, data)
		}
		if err != nil {
			// 对方断开
			if strings.Contains(err.Error(), "Active closure of the user") {
				OnClientDisconected(client)
			}
			// 自己断开
			if strings.Contains(err.Error(), "use of closed network connection") {
				OnServerKicked(client)
			}
		}
	}
}