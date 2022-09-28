// auth: kunlun
// date: 2019-01-16
// description:
package websocket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

// 设置 websocket参数
var upgrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	}, EnableCompression: true,
	HandshakeTimeout: 3 * time.Second,
}

func Handle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Start Handle")
	var (
		wsConn *websocket.Conn
		err    error
		conn   *Connection
		//data   []byte
	)
	if wsConn, err = upgrade.Upgrade(w, r, nil); err != nil {
		return
	}
	if conn, err = InitConnection(wsConn); err != nil {
		goto ERR
	}

	for {

		if _, err = conn.ReadMessage(); err != nil {
			goto ERR
		}

		//if err = conn.WriteMessage(data); err != nil {
		//	goto ERR
		//}

	}

ERR:
	conn.Close()
}
