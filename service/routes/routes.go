package routes

import (
	"fmt"
	"github.com/2dev2/demogo"
	"github.com/2dev2/demogo/service/client"
	"github.com/2dev2/demogo/service/server"
	"log"
	"net/http"
	"os"
)


func SetupRoutes() {

	p := client.NewPool()
	go p.Start()

	//serving the js file as  they are relative to /static folder from root
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(getFileSystem(true))))

	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws",func(w http.ResponseWriter, r *http.Request) {
		server.WsEndpoint(p, w, r)
	} )
}

func getFileSystem(useOS bool) http.FileSystem {
	if useOS {
		log.Print("using live mode of  file system")
		return http.FS(os.DirFS("./web/build/static"))
	}

	log.Print("using embed mode")
	//fsys, err := fs.Sub(demogo.EmbededFiles, "build/static")
	//if err != nil {
	//	panic(err)
	//}

	return http.FS(demogo.EmbededFiles)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	p := "." + r.URL.Path
	if p == "./" {
		p = "./web/build/index.html"
	}
	http.ServeFile(w, r, p)
}
