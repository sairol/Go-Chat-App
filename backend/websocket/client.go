package websocket

import (
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
)

type Client struct {
	Conn *websocket.Conn
	Pool *Pool
	mu   sync.Mutex
}

type Message struct {
	Type int    `json:"type"`
	Body string `json:"body"`
}

func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		messagetype, msg, err := c.Conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		m := Message{Type: messagetype, Body: string(msg)}

		c.Pool.Broadcast <- m

		fmt.Println("msg received = ", m)
	}
}