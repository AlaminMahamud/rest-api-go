package main

import (
	"log"
	"net/http"

	"github.com/alaminmahamud/restApi/api"
)

func main() {
	router := api.New()
	log.Fatal(http.ListenAndServe(":8000", router))
}
