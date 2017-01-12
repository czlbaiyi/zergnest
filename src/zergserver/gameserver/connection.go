package gameserver

import (
	"fmt"

	"golang.org/x/net/websocket"
)

type connection struct {
	// websocket 连接器
	ws *websocket.Conn

	// 发送信息的缓冲 channel
	send chan []byte
}

func (c *connection) reader() {
	var buf []byte
	for {
		if err := websocket.Message.Receive(c.ws, &buf); err != nil {
			break
		} else {
			fmt.Println("从客户端收到：", string(buf))
		}
		h.broadcast <- buf
	}
	c.ws.Close()
}

func (c *connection) writer() {
	for message := range c.send {
		if err := websocket.Message.Send(c.ws, message); err != nil {
			break
		} else {
			fmt.Println("向客户端发送：", string(message))
		}
	}
	c.ws.Close()
}
