package main

import (
	"http-server/middleware"
	"http-server/server"
	"log"
)

func main() {
    srv, err := server.NewServer("127.0.0.1", 8080)
    if err != nil {
        log.Fatal(err)
    }

    // Add middleware
    srv.Use(middleware.Logger)
    srv.Use(middleware.RequestID)
    srv.Use(middleware.CORS)

    log.Printf("Server starting on port 8080")
    if err := srv.Start(); err != nil {
        log.Fatal(err)
    }
}
