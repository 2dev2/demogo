package server

import (
	"fmt"
	"github.com/2dev2/demogo/client"
	"log"
	"net/http"
	"strconv"
)


type Person struct {
	userID string
	message  int
}


var IDCount=0

func WsEndpoint(pool *client.Pool,w http.ResponseWriter, r *http.Request) {
	fmt.Println("WebSocket Endpoint Hit")


	// upgrade this connection to a WebSocket connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	//client connected


	//create web-browser client in our defined format
	//when a new client will join  we will assign a new id and start listening to that client
	IDCount=IDCount+1
	id:=strconv.Itoa(IDCount)
	newClient := client.New(id,ws)

	//store this new connection so that we can reuse later
	pool.Register <- newClient

	// listen indefinitely for new messages coming
	// through on our WebSocket connection
	newClient.Read(pool)

}