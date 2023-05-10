package models

type httpMethod string

const (
	HTTP_GET    httpMethod = "GET"
	HTTP_POST   httpMethod = "POST"
	HTTP_PUT    httpMethod = "PUT"
	HTTP_DELETE httpMethod = "DELETE"
	HTTP_PATCH  httpMethod = "PATCH"
)
