package middleware

import (
	"http-server/server"
	"log"
	"time"
)

// Logger logs request information
func Logger(req *server.Request, res *server.Response, next func() error) error {
    start := time.Now()

    log.Printf("Started %s %s", req.Method, req.Path)

    err := next()

    log.Printf("Completed %s %s in %v", req.Method, req.Path, time.Since(start))

    return err
}

// RequestID adds a request ID header
func RequestID(req *server.Request, res *server.Response, next func() error) error {
    res.SetHeader("X-Request-ID", time.Now().String())
    return next()
}

// CORS adds CORS headers
func CORS(req *server.Request, res *server.Response, next func() error) error {
    res.SetHeader("Access-Control-Allow-Origin", "*")
    res.SetHeader("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
    res.SetHeader("Access-Control-Allow-Headers", "Content-Type")

    if req.Method == "OPTIONS" {
        return nil
    }

    return next()
}
