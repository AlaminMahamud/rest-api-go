package service

import (
	"log"
	"net/http"
)

// StartWebServer - Starts a server in a certain port with some
// error handling code.
func StartWebServer(port string) {

	log.Println("Starting HTTP service at " + port)
	err := http.ListenAndServe(":"+port, nil) // Goroutine will block here

	if err != nil {
		log.Println("An error occured starting HTTP listener at port " + port)
		log.Println("Error: " + err.Error())
	}
}
