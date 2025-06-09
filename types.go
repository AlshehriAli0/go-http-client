package client

import "net/http"

// Route represents a route's pattern, name, and HTTP method.
type Route struct {
	pattern string
	handler HandlerFunc
}

// Middleware defines a function to process middleware logic for a request.
type Middleware func(*Context)

// Group represents a group of routes with a common prefix, useful for modular route organization.
type Group struct {
	prefix string
	app    *App
}

// App is the main application struct for the HTTP client framework.
type App struct {
	routes      map[string]map[Method]Route
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
	Post   Method = "POST"
	Get    Method = "GET"
	Patch  Method = "PATCH"
	Update Method = "UPDATE"
	Delete Method = "delete"
)
