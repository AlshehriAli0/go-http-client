# Go HTTP Client

[![WIP](https://img.shields.io/badge/status-WIP-yellow.svg)](https://github.com/yourusername/go-http-client)
[![GoDoc](https://godoc.org/github.com/AlshehriAli0/go-http-client?status.svg)](https://godoc.org/github.com/AlshehriAli0/go-http-client)
[![Go Report Card](https://goreportcard.com/badge/github.com/AlshehriAli0/go-http-client)](https://goreportcard.com/report/github.com/AlshehriAli0/go-http-client)

A lightweight HTTP client for Go inspired by Express.js patterns. This project aims to provide a familiar development experience for Express.js developers transitioning to Go, while maintaining Go's performance characteristics.

## Installation

### Using Go Modules (Recommended)

```bash
# Initialize your module (if you haven't already)
go mod init your-module-name

# Add the dependency
go get github.com/AlshehriAli0/go-http-client
```

### Manual Installation

```bash
# Clone the repository
git clone https://github.com/AlshehriAli0/go-http-client.git

# Navigate to the project directory
cd go-http-client

# Install the package
go install
```

## Requirements

- Go 1.21 or higher
- Go modules enabled

## Overview

This client provides an Express.js-like interface for building HTTP servers in Go, built on top of the standard `net/http` package. It offers a familiar routing system and middleware support while maintaining Go's performance and simplicity.

## Features

- Express.js style routing
- URL parameter handling (`/:param`)
- Query parameter support
- JSON response handling
- Middleware support
- Minimal API surface
- Standard library based
- Comprehensive error handling
- Cookie support
- Header manipulation
- Status code management

## Documentation

For detailed API documentation, visit [![Go Reference](https://pkg.go.dev/badge/github.com/AlshehriAli0/go-http-client.svg)](https://pkg.go.dev/github.com/AlshehriAli0/go-http-client).

### Key Components

#### App
The main application instance that manages routes and middleware. Use `client.New()` to create a new instance.

#### Context
Provides methods for handling requests and responses, including:
- Request body reading
- Parameter extraction
- Response sending
- Header manipulation
- Cookie handling
- Status code management

#### Middleware
Functions that have access to the request and response objects, and the next middleware function in the application's request-response cycle.

## Examples

### Basic Route
```go
app.Get("/", nil, func(ctx *client.Context) {
    ctx.Send("Hello World")
})
```

### URL Parameters
```go
// Route: "/:id" - Captures any value after the root path as 'id' parameter
// e.g. "/123" -> Returns "123"
app.Get("/:id", nil, func(ctx *client.Context) {
    param := ctx.Param("id")
    ctx.Send(param)
})
```

### Query Parameters
```go
// Route: "/search?s=query" - Retrieves 's' query parameter from URL
// e.g. "/search?s=hello" -> Returns "hello"
app.Get("/search", nil, func(ctx *client.Context) {
    search := ctx.Query("s")
    ctx.Send(search)
})
```

### JSON Response
```go
type Test struct {
    Name string `json:"name"`
    Job  string `json:"job"`
}

app.Get("/json", nil, func(ctx *client.Context) {
    data := Test{Name: "Ali", Job: "dev"}
    ctx.JSON(data)
})
```

### Middleware Example
```go
// Logger middleware prints request information
func Logger(ctx *client.Context) {
    start := time.Now()
    // Process request
    fmt.Printf("Request processed in %v\n", time.Since(start))
}

// Auth middleware example
func AuthMiddleware(ctx *client.Context) {
    token := ctx.GetHeader("Authorization")
    if token == "" {
        ctx.Status(401)
        ctx.JSON(map[string]string{"error": "Unauthorized"})
        ctx.End()
        return
    }
    // Continue to next middleware/handler if authorized
}

// Route with middleware
app.Get("/protected", AuthMiddleware, func(ctx *client.Context) {
    ctx.Send("Protected route - You are authorized!")
})

// Route without middleware (using nil)
app.Get("/public", nil, func(ctx *client.Context) {
    ctx.Send("Public route")
})
```

### Complete Example
```go
package main

import (
    "github.com/AlshehriAli0/go-http-client/client"
)

type Test struct {
    Name string `json:"name"`
    Job  string `json:"job"`
}

func main() {
    app := client.New()
    
    // URL parameter example
    app.Get("/:id", nil, func(ctx *client.Context) {
        param := ctx.Param("id")
        ctx.Send(param)
    })

    // Query parameter example
    app.Get("/search", nil, func(ctx *client.Context) {
        search := ctx.Query("s")
        ctx.Send(search)
    })

    // JSON response example
    app.Get("/json", nil, func(ctx *client.Context) {
        data := Test{Name: "Ali", Job: "dev"}
        ctx.JSON(data)
    })

    // Start the server on port 3000
    app.Start(3000)
}
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## Status

This project is currently a work in progress and is not recommended for production use.

## License

MIT 
