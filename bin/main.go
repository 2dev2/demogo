package main

import (
	"github.com/2dev2/demogo/service/logger"
	"github.com/2dev2/demogo/service/routes"
	"log"
	"net/http"
)

func main(){

	log.SetFlags(0)
	log.SetOutput(logger.LogWriter{})

	routes.SetupRoutes()
	log.Fatal(http.ListenAndServe(":5555", nil))
}
