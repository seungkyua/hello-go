package main

import (
	"log"
	"github.com/gorilla/websocket"
)

type client struct {
	socket	*websocket.Conn
	send	chan []byte
}

func (c *client) read(forward chan<- []byte) {
	defer func() {
		c.socket.Close()
		log.Println("Client socker read closed.")
		close(c.send)
		log.Println("Client send channel closed.")
	}()
	for {
		_, msg, err := c.socket.ReadMessage()
		if err != nil {
			log.Println("Client Socket ReadMaessage error")
			return
		}
		forward <- msg
	}
}

func (c *client) write() {
	go func() {
		defer c.socket.Close()
		for msg := range c.send {
			if err := c.socket.WriteMessage(websocket.TextMessage, msg); err != nil {
				log.Println("Client Socket WriteMaessage error")
				return
			}
		}
	}()
}
