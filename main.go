package main

import (
	"log"
	"os"

	"github.com/ngti/echo_bot/swagger"
)

func main() {

	host := "http://localhost:8080"
	if os.Getenv("EXTENSION_HOST") != "" {
		host = os.Getenv("EXTENSION_HOST")
	}

	apiKey := "9e5cf2779ffd42f88803477b734eb415"
	if os.Getenv("ECHO_API_KEY") != "" {
		apiKey = os.Getenv("ECHO_API_KEY")
	}

	api := swagger.NewDefaultApiWithBasePath(host)
	s := NewServer(&apiKey, api)

	log.Print("Starting server")
	// Binds to $PORT
	s.Run()
}
