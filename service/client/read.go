package client

import (
	"github.com/2dev2/demogo/service/message"
	"log"
)

func (c *Client) Read(p *Pool) {
	defer func() {
		p.Unregister <- c
		c.Conn.Close()
	}()

	for {
		messageType, body, err := c.Conn.ReadMessage()
		log.Print("got the nre Messages from socket")
		if err != nil {
			log.Println(err)
			return
		}

		m := message.Message{Type: messageType,ID:c.ID, Body: string(body)}
		p.Broadcast <- m
		log.Printf("Message Received: %+v\n", m)
	}
}
