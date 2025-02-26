package gale

import (
	"fmt"
	"log"

	"github.com/valyala/fasthttp"
)

type App struct {
	routes map[string]func(Req)
}

func Initialize() *App {
	app := &App{routes: make(map[string]func(Req))}
	return app
}

func (a *App) handler(ctx *fasthttp.RequestCtx) {
	handler, exists := a.routes[string(ctx.Path())]
	if exists {
		handler(Req{context: ctx, body: ctx.PostBody()})
	} else {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		ctx.SetBodyString("Not Found")
	}
}

func (app *App) Listen(args ...int) {
	port := 8070
	if len(args) > 0 {
		port = args[0]
	}
	log.Println("Listening on port", port)
	fasthttp.ListenAndServe(fmt.Sprintf(":%d", port), app.handler)
}
