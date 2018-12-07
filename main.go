package main

import (
	"log"
	"net/http"

	"github.com/alamin-mahamud/rest-api-go/v1/api"
)

func main() {
	router := api.New()
	log.Fatal(http.ListenAndServe(":8000", router))
}
