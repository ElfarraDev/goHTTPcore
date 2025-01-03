package main

import (
	"http-server/server"
	"log"
)

func main() {
	srv, err := server.NewServer("127.0.0.1", 8080)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Server starting on port 8080")
	if err := srv.Start(); err != nil {
		log.Fatal(err)
	}
}
