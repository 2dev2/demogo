package client

import (
	"github.com/2dev2/demogo/service/message"
	"log"
)

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	clients    map[*Client]bool
	Broadcast  chan message.Message
}

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
		Broadcast:  make(chan message.Message),
	}
}

/*
We need to ensure that only one point of our application has the ability to write to
our WebSocket connections or we’ll face concurrent write issues.
So, let’s define our Start() method which will constantly listen for anything passed to
any of our Pool’s channels and then, if anything is received into one of these channels,
it’ll act accordingly.
 */
func (p *Pool) Start() {
	for {
		select {
		case newClient := <-p.Register:
			p.clients[newClient] = true
			log.Println("Size of Connection Pool: ", len(p.clients))
			for c, _ := range p.clients {
				m:=message.Message{Type: 1,ID:newClient.ID, Body: "New User Joined..."}
				log.Println(m)
				c.Conn.WriteJSON(m)
			}
			break
		case leftClient := <-p.Unregister:
			delete(p.clients, leftClient)
			log.Println("Size of Connection Pool: ", len(p.clients))
			for c, _ := range p.clients {
				m:=message.Message{Type: 1, ID:leftClient.ID,Body: "User Disconnected..."}
				c.Conn.WriteJSON(m)
			}
			break
		case m := <-p.Broadcast:
			log.Println("Sending message to all clients in Pool")
			for c, _ := range p.clients {
				if err := c.Conn.WriteJSON(m); err != nil {
					log.Println(err)
					return
				}
			}
		}
	}
}