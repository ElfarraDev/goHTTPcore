package server

type Request struct {
	Method      string            // GET, POST, PUT, DELETE
	Path        string            // /api/v1/users or /api/v1/users/1 or /api/v1/users/1?name=John
	Version     string            // HTTP/1.1 or HTTP/2 or HTTP/3
	Headers     map[string]string // metadata about the request (content type, content length, etc.)
	Body        []byte            // request body: json data or form data or file data
	QueryParams map[string]string // query parameters: name=John&age=30
	Remoteaddr  string            // remote address of the client
	KeepAlive   bool              // keep the connection alive
}
