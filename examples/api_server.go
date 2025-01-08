package examples

import (
	"encoding/json"
	"http-server/middleware"
	"http-server/server"
	"log"
	"time"
)

// SimpleAuthMiddleware checks for a basic API key
func SimpleAuthMiddleware(req *server.Request, res *server.Response, next func() error) error {
	apiKey := req.Headers["X-API-Key"]
	if apiKey != "your-secret-key" {
		res.SetStatus(401)
		return res.Write([]byte("Unauthorized"))
	}
	return next()
}

// RateLimiterMiddleware demonstrates a basic rate limiter
func RateLimiterMiddleware(req *server.Request, res *server.Response, next func() error) error {
	// For demo purposes, limit to 1 request per second based on IP
	if isRateLimited(req.Remoteaddr) {
		res.SetStatus(429)
		return res.Write([]byte("Too Many Requests"))
	}
	return next()
}

func RunAPIServer() {
	srv, err := server.NewServer("127.0.0.1", 8080)
	if err != nil {
		log.Fatal(err)
	}

	// Global middleware
	srv.Use(middleware.Logger)
	srv.Use(middleware.CORS)
	srv.Use(RateLimiterMiddleware)

	// Protected endpoints require auth
	srv.Use(SimpleAuthMiddleware)

	// Example handlers
	srv.Handle("GET", "/api/users", handleGetUsers)
	srv.Handle("POST", "/api/users", handleCreateUser)
	srv.Handle("GET", "/healthcheck", handleHealthCheck)

	log.Printf("API Server starting on port 8080")
	if err := srv.Start(); err != nil {
		log.Fatal(err)
	}
}

// Example handlers
func handleGetUsers(req *server.Request, res *server.Response) error {
	users := []map[string]string{
		{"id": "1", "name": "John Doe"},
		{"id": "2", "name": "Jane Smith"},
	}

	jsonData, err := json.Marshal(users)
	if err != nil {
		return err
	}

	res.SetHeader("Content-Type", "application/json")
	return res.Write(jsonData)
}

func handleCreateUser(req *server.Request, res *server.Response) error {
	var newUser map[string]string
	if err := json.Unmarshal(req.Body, &newUser); err != nil {
		res.SetStatus(400)
		return res.Write([]byte("Invalid JSON"))
	}

	// Mock successful creation
	res.SetStatus(201)
	res.SetHeader("Content-Type", "application/json")
	return res.Write(req.Body)
}

func handleHealthCheck(req *server.Request, res *server.Response) error {
	health := map[string]string{
		"status": "healthy",
		"time":   time.Now().String(),
	}

	jsonData, err := json.Marshal(health)
	if err != nil {
		return err
	}

	res.SetHeader("Content-Type", "application/json")
	return res.Write(jsonData)
}

// Helper function for rate limiting
var lastRequestTime = make(map[string]time.Time)

func isRateLimited(ip string) bool {
	lastTime, exists := lastRequestTime[ip]
	now := time.Now()

	if !exists || now.Sub(lastTime) > time.Second {
		lastRequestTime[ip] = now
		return false
	}
	return true
}
