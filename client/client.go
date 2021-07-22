package client

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	ID   string
	Conn *websocket.Conn
	//Pool *Pool
}

func New(ID   string,Conn *websocket.Conn)*Client{
	return &Client{
		ID,
		Conn,
	}
}



