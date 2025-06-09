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

This client provides an Express.js like interface for building HTTP servers in Go, built on top of the standard `net/http` package. It offers a familiar routing system and middleware support while maintaining Go's performance and simplicity.

**It is extremely easy to use and lightweight, specifically designed to help Express.js (Node.js) developers transition to Go with minimal friction.**

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
- Supports multiple routes with the same name but different HTTP methods (e.g., both `GET /user` and `POST /user`) unlike http/net package

## Comparison: Our API vs. Writing Go HTTP Servers from Scratch

When building HTTP servers in Go using only the standard library, you typically need to:
- Manually set up route matching and parsing
- Write boilerplate for extracting URL/query parameters
- Chain middleware logic by hand
- Manage context, headers, cookies, and status codes explicitly
- Handle JSON serialization and error responses yourself

With **go-http-client**, you get:
- Express.js style route registration with method based handlers (e.g., `app.Get`, `app.Post`)
- Built in support for multiple routes with the same name but different HTTP methods
- Automatic parameter and query extraction via `ctx.Param` and `ctx.Query`
- Simple middleware chaining with `app.Use` and per route middleware
- Unified `Context` object for request/response handling
- Built in helpers for JSON, status codes, headers, and cookies
- Less boilerplate and more readable, maintainable code

**Example: Registering multiple methods for the same route**

```go
app.Get("/user", nil, func(ctx *client.Context) {
    ctx.Send("GET user")
})
app.Post("/user", nil, func(ctx *client.Context) {
    ctx.Send("POST user")
})
```

**Standard library equivalent:**

```go
http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        w.Write([]byte("GET user"))
    case "POST":
        w.Write([]byte("POST user"))
    default:
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
    }
})
```

**Example: Dynamic Routing with Params and Middleware**

**With go-http-client:**

```go
// Logger middleware
func Logger(ctx *client.Context) {
    fmt.Printf("%s %s\n", ctx.Method(), ctx.Path())
}

app := client.New()
app.Use(Logger)

// Dynamic route with param and per-route middleware
auth := func(ctx *client.Context) {
    if ctx.GetHeader("Authorization") == "" {
        ctx.Status(401)
        ctx.Send("Unauthorized")
        ctx.End()
    }
}

app.Get("/user/:id", auth, func(ctx *client.Context) {
    ctx.Send("GET user " + ctx.Param("id"))
})
app.Post("/user/:id", auth, func(ctx *client.Context) {
    ctx.Send("POST user " + ctx.Param("id"))
})

app.Start(3000)
```

**Standard library equivalent:**

```go
http.HandleFunc("/user/", func(w http.ResponseWriter, r *http.Request) {
    // Logger
    fmt.Printf("%s %s\n", r.Method, r.URL.Path)

    // Basic param extraction
    parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
    if len(parts) < 2 || parts[0] != "user" {
        http.NotFound(w, r)
        return
    }
    id := parts[1]

    // Auth middleware
    if r.Header.Get("Authorization") == "" {
        w.WriteHeader(401)
        w.Write([]byte("Unauthorized"))
        return
    }

    switch r.Method {
    case "GET":
        w.Write([]byte("GET user " + id))
    case "POST":
        w.Write([]byte("POST user " + id))
    default:
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
    }
})

http.ListenAndServe(":3000", nil)
```

With **go-http-client**, you get expressive, readable, and maintainable code for dynamic routes, params, and middleware—without boilerplate or manual parsing.

**Another Example: Same Power, Half the Code**

**With go-http-client:**

```go
// Logger middleware
func Logger(ctx *client.Context) {
    fmt.Printf("%s %s\n", ctx.Method(), ctx.Path())
}

// Auth middleware
func Auth(ctx *client.Context) {
    if ctx.GetHeader("X-Token") != "secret" {
        ctx.Status(401)
        ctx.JSON(map[string]string{"error": "unauthorized"})
        ctx.End()
    }
}

app := client.New()
app.Use(Logger)

app.Get("/greet/:name", Auth, func(ctx *client.Context) {
    lang := ctx.Query("lang")
    name := ctx.Param("name")
    msg := "Hello, " + name
    if lang == "es" {
        msg = "Hola, " + name
    }
    ctx.JSON(map[string]string{"message": msg})
})

app.Post("/greet/:name", Auth, func(ctx *client.Context) {
    ctx.JSON(map[string]string{"message": "Posted to " + ctx.Param("name")})
})

app.Start(3000)
```

**Standard library equivalent:**

```go
http.HandleFunc("/greet/", func(w http.ResponseWriter, r *http.Request) {
    // Logger
    fmt.Printf("%s %s\n", r.Method, r.URL.Path)

    // Auth
    if r.Header.Get("X-Token") != "secret" {
        w.WriteHeader(401)
        w.Header().Set("Content-Type", "application/json")
        w.Write([]byte(`{"error":"unauthorized"}`))
        return
    }

    // Param extraction
    parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
    if len(parts) < 2 || parts[0] != "greet" {
        http.NotFound(w, r)
        return
    }
    name := parts[1]

    switch r.Method {
    case "GET":
        lang := r.URL.Query().Get("lang")
        msg := "Hello, " + name
        if lang == "es" {
            msg = "Hola, " + name
        }
        w.Header().Set("Content-Type", "application/json")
        w.Write([]byte(fmt.Sprintf(`{"message":"%s"}`, msg)))
    case "POST":
        w.Header().Set("Content-Type", "application/json")
        w.Write([]byte(fmt.Sprintf(`{"message":"Posted to %s"}", name)))
    default:
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
    }
})

http.ListenAndServe(":3000", nil)
```

With **go-http-client**, you achieve the same result with about half the code, no manual parsing, and much greater clarity.

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
Functions that have access to the request and response objects, and the next middleware function in the application's request response cycle.

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
        ctx.End() // using end will cancel middleware chaining so maybe redirect if this isn't intended
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

### Route Grouping for Modular Systems

Route grouping allows you to organize related routes under a common prefix, making your codebase more modular and maintainable. Each group can have its own middleware and handlers, and is ideal for separating concerns (e.g., user routes, admin routes).

```go
app := client.New()

// Global middleware
app.Use(func(ctx *client.Context) {
    fmt.Printf("[LOG] %s %s\n", ctx.Method(), ctx.Path())
})

// Group for user-related routes
userGroup := app.Group("/users")

// Now all routes will start with /users
userGroup.Get("/", nil, func(ctx *client.Context) {
    ctx.Send("List all users")
})
userGroup.Get("/:id", nil, func(ctx *client.Context) {
    ctx.Send("Get user with ID: " + ctx.Param("id"))
})
userGroup.Post("/", AuthMiddleware, func(ctx *client.Context) {
    ctx.Send("Create user (auth required)")
})

// Group for admin-related routes
adminGroup := app.Group("/admin")
adminGroup.Get("/dashboard", AuthMiddleware, func(ctx *client.Context) {
    ctx.Send("Welcome to the admin dashboard!")
})

app.Start(3000)
```

This makes it easy to keep your route logic modular and organized, especially for larger applications.

## Status

This project is currently a work in progress and is not recommended for production use yet.

## License

MIT 
