package server

import (
	"http-server/middleware"
	"log"
	"net"
)

type Server struct {
    socket     *Socket
    listener   net.Listener
    middleware *middleware.Chain
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
        middleware: middleware.NewChain(),
    }, nil
}

func (s *Server) Use(handler middleware.Handler) {
    s.middleware.Use(handler)
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
