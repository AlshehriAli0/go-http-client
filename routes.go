package client

// Default routes

// Get registers a new GET route with the given path and handler function.
func (app *App) Get(route string, mw Middleware, handler HandlerFunc) {
	normalizedRoute := normalizePath(route)
	app.handle(Get, normalizedRoute, mwToSlice(mw), handler)
}

// Post registers a new POST route with the given path and handler function.
func (app *App) Post(route string, mw Middleware, handler HandlerFunc) {
	normalizedRoute := normalizePath(route)
	app.handle(Post, normalizedRoute, mwToSlice(mw), handler)
}

// Update registers a new UPDATE route with the given path and handler function.
func (app *App) Update(route string, mw Middleware, handler HandlerFunc) {
	normalizedRoute := normalizePath(route)
	app.handle(Update, normalizedRoute, mwToSlice(mw), handler)
}

// Patch registers a new PATCH route with the given path and handler function.
func (app *App) Patch(route string, mw Middleware, handler HandlerFunc) {
	normalizedRoute := normalizePath(route)
	app.handle(Patch, normalizedRoute, mwToSlice(mw), handler)
}

// Delete registers a new DELETE route with the given path and handler function.
func (app *App) Delete(route string, mw Middleware, handler HandlerFunc) {
	normalizedRoute := normalizePath(route)
	app.handle(Delete, normalizedRoute, mwToSlice(mw), handler)
}

// Prefixed Routes

// Get registers a new GET route with the given path and handler function for the group.
func (g *Group) Get(route string, mw Middleware, handler HandlerFunc) {
	fullRoute := g.prefix + route
	normalizedRoute := normalizePath(fullRoute)
	mws := combineMiddleware(g.middlewares, mw)
	g.app.handle(Get, normalizedRoute, mws, handler)
}

// Post registers a new POST route with the given path and handler function for the group.
func (g *Group) Post(route string, mw Middleware, handler HandlerFunc) {
	fullRoute := g.prefix + route
	normalizedRoute := normalizePath(fullRoute)
	mws := combineMiddleware(g.middlewares, mw)
	g.app.handle(Post, normalizedRoute, mws, handler)
}

// Update registers a new UPDATE route with the given path and handler function for the group.
func (g *Group) Update(route string, mw Middleware, handler HandlerFunc) {
	fullRoute := g.prefix + route
	normalizedRoute := normalizePath(fullRoute)
	mws := combineMiddleware(g.middlewares, mw)

	g.app.handle(Update, normalizedRoute, mws, handler)
}

// Patch registers a new PATCH route with the given path and handler function for the group.
func (g *Group) Patch(route string, mw Middleware, handler HandlerFunc) {
	fullRoute := g.prefix + route
	normalizedRoute := normalizePath(fullRoute)
	mws := combineMiddleware(g.middlewares, mw)

	g.app.handle(Patch, normalizedRoute, mws, handler)
}

// Delete registers a new DELETE route with the given path and handler function for the group.
func (g *Group) Delete(route string, mw Middleware, handler HandlerFunc) {
	fullRoute := g.prefix + route
	normalizedRoute := normalizePath(fullRoute)
	mws := combineMiddleware(g.middlewares, mw)

	g.app.handle(Delete, normalizedRoute, mws, handler)
}
