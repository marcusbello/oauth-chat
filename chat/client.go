package main

import "github.com/gorilla/websocket"

// client represents single chat client
type client struct {
	//socket is the web socket
	socket *websocket.Conn
	//send is the channel on which message are sent
	send chan []byte
	//room is the room this client is chatting
	room *room
}

func (c *client) read() {
	defer c.socket.Close()
	for {
		_, msg, err := c.socket.ReadMessage()
		if err != nil {
			return
		}
		c.room.forward <- msg
	}
}

func (c *client) write() {
	defer c.socket.Close()

	for msg := range c.send {
		err := c.socket.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			return
		}
	}
}
