package middleware

import (
	"http-server/server"
)

// Handler represents a middleware function
type Handler func(*server.Request, *server.Response, func() error) error

// Chain represents a chain of middleware handlers
type Chain struct {
    handlers []Handler
}

// NewChain creates a new middleware chain
func NewChain() *Chain {
    return &Chain{
        handlers: make([]Handler, 0),
    }
}

// Use adds a middleware handler to the chain
func (c *Chain) Use(handler Handler) {
    c.handlers = append(c.handlers, handler)
}

// Execute runs all middleware in the chain
func (c *Chain) Execute(req *server.Request, res *server.Response, final func() error) error {
    var (
        index    = 0
        handlers = c.handlers
    )

    var next func() error
    next = func() error {
        if index >= len(handlers) {
            return final()
        }
        index++
        return handlers[index-1](req, res, next)
    }

    return next()
}
