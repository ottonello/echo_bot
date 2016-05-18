package main

import "log"

func main() {
	s := NewServer()

	log.Print("Starting server")
	// Binds to $PORT
	s.Run()
}
