package response

import (
	"io"
)

type Response struct {
	StatusCode int               // 200, 404, 500, etc.
	StatusText string            // Text description of the status code
	Headers    map[string]string // metadata about the response (content type, content length, etc.)
	Body       []byte            // response body: json data or form data or file data
	Write      io.Writer         // write the response to the client
}

var StatusCodes = map[int]string{
	// 2xx Success
	200: "OK",
	201: "Created",
	202: "Accepted",
	204: "No Content",

	// 3xx Redirection
	301: "Moved Permanently",
	302: "Found",
	303: "See Other",
	304: "Not Modified",

	// 4xx Client Error
	400: "Bad Request",
	401: "Unauthorized",
	403: "Forbidden",
	404: "Not Found",
	429: "Too Many Requests",

	// 5xx Server Error
	500: "Internal Server Error",
	501: "Not Implemented",
	502: "Bad Gateway",
	503: "Service Unavailable",
}

// NewResponse creates a new Response instance
func NewResponse(writer io.Writer) *Response {
	return &Response{
		Headers:    make(map[string]string),
		StatusCode: 200,
		StatusText: StatusCodes[200],
		Writer:     writer,
	}
}

// SetStatus sets the response status code and text
func (r *Response) SetStatus(code int) {
	r.StatusCode = code
	if text, ok := StatusCodes[code]; ok {
		r.StatusText = text
	} else {
		r.StatusText = "Unknown Status"
	}
}

// SetHeader sets a response header
func (r *Response) SetHeader(key, value string) {
	r.Headers[key] = value
}

// Write writes data to the response body
func (r *Response) Write(data []byte) error {
	r.Body = append(r.Body, data...)
	return nil
}
