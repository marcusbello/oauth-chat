package main

import (
	"github.com/gorilla/websocket"
	"time"
)

// client represents single chat client
type client struct {
	//socket is the web socket
	socket *websocket.Conn
	//send is the channel on which message are sent
	send chan *message
	//room is the room this client is chatting
	room *room
	//userData holds info about user
	userData map[string]interface{}
}

func (c *client) read() {
	defer c.socket.Close()
	for {
		var msg *message
		err := c.socket.ReadJSON(&msg)
		if err != nil {
			return
		}
		msg.When = time.Now().UTC()
		msg.Name = c.userData["name"].(string)
		c.room.forward <- msg
	}
}

func (c *client) write() {
	defer c.socket.Close()

	for msg := range c.send {
		err := c.socket.WriteJSON(msg)
		if err != nil {
			return
		}
	}
}
