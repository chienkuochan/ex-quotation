// auth: kunlun
// date: 2019-01-16
// description:
package request

import (
	"github.com/gorilla/websocket"
)

type Yihuo interface {
	// 心跳
	Pong()
	// 连接
	WsConnect()
	// 订阅
	Subscribe()
	//读取消息
	ReadMessage()
	// 成交
	YiHuoTickWebsocket()
	// 深度
	YiHuoDepthWebSocket()
}

type yihuo struct {
	URL string
	Ws  *websocket.Conn
}
