package main

import (
	"log"

	"github.com/ALEJORIOS/gale"
)

func main() {
	app := gale.Initialize()
	app.Get("/test", testHandler)
	app.Post("/test", testHandler)
	app.Listen()
}

func testHandler(req gale.Req) {
	json, _ := req.JsonBody()
	log.Println("Hello, this:", json)
	req.Send("Esta es mi respuesta desde el servidor")
}
