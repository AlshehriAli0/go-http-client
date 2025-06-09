package client

import (
	"encoding/json"
	"io"
	"net/http"
)

// End marks the context as terminated, stopping further middleware execution.
func (ctx *Context) End() {
	ctx.terminated = true
}

// ReadBody reads and returns the request body as a string.
func (ctx *Context) ReadBody() (string, error) {
	defer ctx.Request.Body.Close()

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// Redirect performs an HTTP redirect to the specified route with the given status code.
func (ctx *Context) Redirect(route string, code int) {
	http.Redirect(ctx.Writer, ctx.Request, route, code)
	ctx.End()
}

// Send writes a string response to the client.
func (ctx *Context) Send(str string) {
	strByte := []byte(str)
	ctx.Writer.Write(strByte)
}

// Error sends an HTTP error response with the given message and status code.
func (ctx *Context) Error(message string, code int) {
	http.Error(ctx.Writer, message, code)
}

// Param retrieves the value of a URL parameter by its key.
// Example: For route "/users/:id", Param("id") returns the actual id value.
func (ctx *Context) Param(paramKey string) string {
	return ctx.param[paramKey]
}

// Query retrieves the value of a query parameter from the URL.
// Example: For URL "/search?q=test", Query("q") returns "test".
func (ctx *Context) Query(queryParam string) string {
	query := ctx.Request.URL.Query()
	return query.Get(queryParam)
}

// JSON marshals the provided data into JSON and sends it as response.
// Sets appropriate Content-Type header automatically.
func (ctx *Context) JSON(jsonData interface{}) {
	parsedJson, err := json.Marshal(jsonData)

	if err != nil {
		ctx.Error("Not a valid JSON", http.StatusInternalServerError)
		return
	}

	ctx.Writer.Header().Set("Content-Type", "application/json")
	ctx.Writer.Write(parsedJson)
}

// Status sets the HTTP status code for the response.
// This should be called before writing any response body.
func (ctx *Context) Status(code int) {
	ctx.Writer.WriteHeader(code)
}

// SetHeader sets a response header key to a given value.
// If the header already exists, it will be overwritten.
func (ctx *Context) SetHeader(key, value string) {
	ctx.Writer.Header().Set(key, value)
}

// AddHeader appends a value to a header key (without overwriting existing values).
func (ctx *Context) AddHeader(key, value string) {
	ctx.Writer.Header().Add(key, value)
}

// GetHeader retrieves a header value from the incoming request.
// Returns an empty string if the header is not set.
func (ctx *Context) GetHeader(key string) string {
	return ctx.Request.Header.Get(key)
}

// SendStatus sends only an HTTP status code and its corresponding status text.
// Useful for quick status replies like 404, 403, etc.
func (ctx *Context) SendStatus(code int) {
	ctx.Writer.WriteHeader(code)
	ctx.Writer.Write([]byte(http.StatusText(code)))
}

// SetCookie sets an HTTP cookie on the response.
func (ctx *Context) SetCookie(cookie *http.Cookie) {
	http.SetCookie(ctx.Writer, cookie)
}

// Cookie retrieves a named cookie from the request.
// Returns the cookie and an error if it's not found or malformed.
func (ctx *Context) Cookie(name string) (*http.Cookie, error) {
	return ctx.Request.Cookie(name)
}
