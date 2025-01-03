package server

import (
	"log"
	"net"
)

type Server struct {
	socket   *Socket
	listener net.Listener
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
		listener: listener,
	}, nil
}

func (s *Server) Start() error {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			continue
		}

		go s.handleConnection(conn)
	}
}

func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()

	handler := NewHTTPHandler(conn)
	err := handler.HandleRequest()
	if err != nil {
		log.Printf("Error handling request: %v", err)
	}
}
