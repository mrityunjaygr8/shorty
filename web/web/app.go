package web

import (
	"github.com/gorilla/mux"
	"github.com/mrityunjaygr8/shorty/app"
)

type WebApp struct {
	host   string
	port   uint
	app    app.App
	router *mux.Router
}
