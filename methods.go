package gale

func (a *App) Get(path string, handler func(Req)) {
	a.addRoute(path, handler)
}

func (a *App) Post(path string, handler func(Req)) {

	a.addRoute(path, handler)
}

func (a *App) addRoute(path string, handler func(Req)) {
	a.routes[path] = handler
}
