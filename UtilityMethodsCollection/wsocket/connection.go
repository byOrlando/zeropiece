package wsocket

import (
	"errors"
	"github.com/gorilla/websocket"
	"sync"
)

type Connection struct {
	WsConn    *websocket.Conn
	InChan    chan []byte
	OutChan   chan []byte
	CloseChan chan []byte
	Muter     sync.Mutex
	IsClosed  bool
}

func InitConnection(wsConn *websocket.Conn) (conn *Connection, err error) {
	conn = &Connection{
		WsConn:    wsConn,
		InChan:    make(chan []byte, 1000),
		OutChan:   make(chan []byte, 1000),
		CloseChan: make(chan []byte, 1),
	}
	//启动读协程
	go conn.readLoop()
	//启动写协程
	go conn.writeLoop()
	return
}

func (conn *Connection) ReadMessage() (data []byte, err error) {
	select {
	case data = <-conn.InChan:
	case <-conn.CloseChan:
		err = errors.New("connection is closed")
	}

	return
}

func (conn *Connection) WriteMessage(data []byte) (err error) {
	select {
	case conn.OutChan <- data:
	case <-conn.CloseChan:
		err = errors.New("connection is closed")
	}
	return
}

func (conn *Connection) Close() {
	//线程安全
	conn.WsConn.Close()

	//以下代码执行一次
	conn.Muter.Lock()
	if !conn.IsClosed {
		close(conn.CloseChan)
		conn.IsClosed = true
	}
	conn.Muter.Unlock()
}

func (conn *Connection) readLoop() {

	for {
		select {
		case data := <-conn.InChan:
			if string(data) == "\"ping\"" || string(data) == "ping" {
				conn.OutChan <- []byte("pong")
			} else {
				conn.MessageHandler(data)
			}
		case <-conn.CloseChan:
			//closeChan关闭的时候
			goto ERR

		}

	}
ERR:
	conn.Close()

}

//
func (conn *Connection) writeLoop() {
	var (
		err error
		str []byte
	)
	for {
		select {
		case data := <-conn.OutChan:
			if string(data) == "pong" {
				str = []byte("pong")
			} else {
				str = data
			}

		case <-conn.CloseChan:
			goto ERR
		}

		if err = conn.WsConn.WriteMessage(websocket.TextMessage, str); err != nil {
			goto ERR
		}
	}
ERR:
	conn.Close()
}
