package http

type HTTPHandler struct {
	request  *Request
	response *Response
}

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net"
	"strconv"
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
