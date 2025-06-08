package client

import "net/http"

// RouteEntry represents a route's pattern, name, and HTTP method.
type RouteEntry struct {
	pattern   string
	routeName string
	method    Method
}

// Middleware defines a function to process middleware logic for a request.
type Middleware func(*Context)

// App is the main application struct for the HTTP client framework.
type App struct {
	routes      map[RouteEntry]HandlerFunc
	middlewares []Middleware
}

// HandlerFunc defines the handler used for processing HTTP requests.
type HandlerFunc func(*Context)

// Context holds information about the current HTTP request and response.
type Context struct {
	Writer     http.ResponseWriter
	Request    *http.Request
	param      map[string]string
	terminated bool
}

// Method represents an HTTP method type.
type Method string

const (
	// Post represents the HTTP POST method.
	Post Method = "post"
	// Get represents the HTTP GET method.
	Get Method = "get"
	// Patch represents the HTTP PATCH method.
	Patch Method = "patch"
	// Update represents the HTTP UPDATE method.
	Update Method = "update"
	// Delete represents the HTTP DELETE method.
	Delete Method = "delete"
)
