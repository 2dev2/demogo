package main

import (
	"fmt"
	"github.com/2dev2/demogo/routes"
	"github.com/2dev2/demogo/service"
	"log"
	"net/http"
)

func main(){


	routes.SetupRoutes()
	svc:=service.NewSvc()
	fmt.Print(svc)
	fmt.Print("dev")
	log.Fatal(http.ListenAndServe(":5555", nil))
}
