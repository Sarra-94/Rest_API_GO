package main

import (
	"goTutorial/Helpers"
	"goTutorial/Routers"
	"log"
	"net/http"
)

func main() {
	// Init router
	Helpers.Migration()

	// Start server
	r:= Routers.InitRouter()
	log.Fatal(http.ListenAndServe(":1230", r))
}

