package wsocket

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
)

//读写锁
var rwlocker sync.RWMutex

// WsHandler ws链接处理
func WsHandler(c *gin.Context) {
	var (
		wsConn *websocket.Conn
		err    error
		conn   *Connection
		data   []byte
	)
	userName := c.Query("id")
	fmt.Println("有新的链接--userName:", userName)

	if userName == "" {
		c.Abort()
		return
	}
	if wsConn, err = (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		Subprotocols: []string{c.Request.Header.Get("Sec-WebSocket-Protocol")},
	}).Upgrade(c.Writer, c.Request, nil); err != nil {
		return
	}

	if conn, err = InitConnection(wsConn); err != nil {
		fmt.Println("InitConnection err:", err)
		goto ERR
	}
	rwlocker.Lock()
	WSSET(userName, conn)
	rwlocker.Unlock()
	for {
		if _, data, err = conn.WsConn.ReadMessage(); err != nil {
			goto ERR
		}
		conn.InChan <- data
	}
ERR:
	//关闭链接操作
	conn.Close()
	//删除映射
	rwlocker.Lock()
	DELWS(userName, conn)
	rwlocker.Unlock()

}
