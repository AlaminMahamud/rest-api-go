package main

import (
	"fmt"
	"os"

	"github.com/alamin-mahamud/rest-api-go/v2/services"
)

var appName = "accountservice"

func main() {
	port := os.Getenv("port")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("Starting %s\n", appName)
	service.StartWebServer(port)
}
