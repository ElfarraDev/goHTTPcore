package server

import (
	"log"
	"net"
)

type Server struct {
	socket     *Socket
	listener   net.Listener
	middleware *Chain
	routes     map[string]map[string]func(*Request, *Response) error
}

func NewServer(ip string, port uint16) (*Server, error) {
	addr := &net.TCPAddr{
		IP:   net.ParseIP(ip),
		Port: int(port),
	}

	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return nil, err
	}

	return &Server{
		socket: &Socket{
			IPAdress: addr.IP,
			Port:     port,
		},
		listener:   listener,
		middleware: NewChain(),
		routes:     make(map[string]map[string]func(*Request, *Response) error),
	}, nil
}

func (s *Server) Use(handler Handler) {
	s.middleware.Use(handler)
}

// Handle registers a route handler for the given method and path.
func (s *Server) Handle(method, path string, handler func(*Request, *Response) error) {
	if s.routes[method] == nil {
		s.routes[method] = make(map[string]func(*Request, *Response) error)
	}
	s.routes[method][path] = handler
}

// Start begins accepting connections and handling requests.
func (s *Server) Start() error {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			// Stop the server if the listener is closed.
			if ne, ok := err.(net.Error); ok && !ne.Temporary() {
				return err
			}
			log.Printf("Accept error: %v", err)
			continue
		}
		go s.handleConnection(conn)
	}
}

func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()

	handler := NewHTTPHandler(conn)

	err := s.middleware.Execute(handler.request, handler.response, func() error {
		return handler.HandleRequest()
	})

	if err != nil {
		log.Printf("Error handling request: %v", err)
	}
}
