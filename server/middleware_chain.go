package server

// Handler represents a middleware function
// It receives a request and response and calls the next handler in the chain.
type Handler func(*Request, *Response, func() error) error

// Chain represents a chain of middleware handlers
// allowing sequential execution of middleware.
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

// Execute runs all middleware in the chain then calls final
func (c *Chain) Execute(req *Request, res *Response, final func() error) error {
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
