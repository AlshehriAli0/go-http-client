package client

// Default routes

// Get registers a new GET route with the given path and handler function.
func (app *App) Get(route string, mw Middleware, handler HandlerFunc) {
	normalizedRoute := normalizePath(route)
	app.handle(Get, normalizedRoute, mw, handler)
}

// Post registers a new POST route with the given path and handler function.
func (app *App) Post(route string, mw Middleware, handler HandlerFunc) {
	normalizedRoute := normalizePath(route)
	app.handle(Post, normalizedRoute, mw, handler)
}

// Update registers a new UPDATE route with the given path and handler function.
func (app *App) Update(route string, mw Middleware, handler HandlerFunc) {
	normalizedRoute := normalizePath(route)
	app.handle(Update, normalizedRoute, mw, handler)
}

// Patch registers a new PATCH route with the given path and handler function.
func (app *App) Patch(route string, mw Middleware, handler HandlerFunc) {
	normalizedRoute := normalizePath(route)
	app.handle(Patch, normalizedRoute, mw, handler)
}

// Delete registers a new DELETE route with the given path and handler function.
func (app *App) Delete(route string, mw Middleware, handler HandlerFunc) {
	normalizedRoute := normalizePath(route)
	app.handle(Delete, normalizedRoute, mw, handler)
}

// Prefixed Routes

// Get registers a new GET route with the given path and handler function for the group.
func (g *Group) Get(route string, mw Middleware, handler HandlerFunc) {
	fullRoute := g.prefix + route
	normalizedRoute := normalizePath(fullRoute)
	g.app.handle(Get, normalizedRoute, mw, handler)
}

// Post registers a new POST route with the given path and handler function for the group.
func (g *Group) Post(route string, mw Middleware, handler HandlerFunc) {
	fullRoute := g.prefix + route
	normalizedRoute := normalizePath(fullRoute)
	g.app.handle(Post, normalizedRoute, mw, handler)
}

// Update registers a new UPDATE route with the given path and handler function for the group.
func (g *Group) Update(route string, mw Middleware, handler HandlerFunc) {
	fullRoute := g.prefix + route
	normalizedRoute := normalizePath(fullRoute)

	g.app.handle(Update, normalizedRoute, mw, handler)
}

// Patch registers a new PATCH route with the given path and handler function for the group.
func (g *Group) Patch(route string, mw Middleware, handler HandlerFunc) {
	fullRoute := g.prefix + route
	normalizedRoute := normalizePath(fullRoute)

	g.app.handle(Patch, normalizedRoute, mw, handler)
}

// Delete registers a new DELETE route with the given path and handler function for the group.
func (g *Group) Delete(route string, mw Middleware, handler HandlerFunc) {
	fullRoute := g.prefix + route
	normalizedRoute := normalizePath(fullRoute)

	g.app.handle(Delete, normalizedRoute, mw, handler)
}
