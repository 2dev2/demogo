package users

import "github.com/gorilla/websocket"

var AllUsers []*websocket.Conn


func contains(s []*websocket.Conn, e *websocket.Conn) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}