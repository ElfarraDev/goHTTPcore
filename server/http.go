package server

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

// HTTPHandler manages the HTTP request/response cycle
type HTTPHandler struct {
    request  *Request
    response *Response
    conn     net.Conn
}

// NewHTTPHandler creates a new HTTP handler for a connection
func NewHTTPHandler(conn net.Conn) *HTTPHandler {
    return &HTTPHandler{
        conn:     conn,
        request:  &Request{},
        response: NewResponse(conn),
    }
}

func (h *HTTPHandler) HandleRequest() error {
    reader := bufio.NewReader(h.conn)

    // Read request line
    requestLine, err := reader.ReadString('\n')
    if err != nil {
        return err
    }

    // Parse request line
    parts := strings.Split(strings.TrimSpace(requestLine), " ")
    if len(parts) != 3 {
        return fmt.Errorf("invalid request line")
    }

    h.request.Method = parts[0]
    h.request.Path = parts[1]
    h.request.Version = parts[2]

    // Simple response for testing
    response := fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\n\r\nRequested Path: %s\n", h.request.Path)
    h.conn.Write([]byte(response))

    return nil
}
