package main

import (
	"fmt"
	_ "net/http/pprof"
	"os"

	"github.com/alamin-mahamud/rest-api-go/v3/services"
)

var appName = "accountservice"

func main() {
	port := os.Getenv("port")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("Starting %s\n", appName)
	services.StartWebServer(port)
}
