package examples

import (
	"http-server/middleware"
	"http-server/server"
	"log"
)

func RunSimpleServer() {
	srv, err := server.NewServer("127.0.0.1", 8080)
	if err != nil {
		log.Fatal(err)
	}

	srv.Use(middleware.Logger)    // Log all requests
	srv.Use(middleware.RequestID) // Add request IDs

	srv.Use(func(req *server.Request, res *server.Response, next func() error) error {
		res.SetHeader("X-Custom-Header", "Hello from custom middleware!")
		return next()
	})

	log.Printf("Example server starting on port 8080")
	if err := srv.Start(); err != nil {
		log.Fatal(err)
	}
}
