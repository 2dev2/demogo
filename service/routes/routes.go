package routes

import (
	"fmt"
	"github.com/2dev2/demogo/service/client"
	"github.com/2dev2/demogo/service/server"
	"net/http"
)

func SetupRoutes() {

	p := client.NewPool()
	go p.Start()

	//serving the js file as  they are relative to /static folder from root
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./web/build/static"))))

	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws",func(w http.ResponseWriter, r *http.Request) {
		server.WsEndpoint(p, w, r)
	} )
}


func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	p := "." + r.URL.Path
	if p == "./" {
		p = "./web/build/index.html"
	}
	http.ServeFile(w, r, p)
}
