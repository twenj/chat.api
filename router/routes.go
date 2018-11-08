package router

import (
	"github.com/gorilla/mux"
	"go-web/app"
	"goinsurance/handler"
)

func Init() *mux.Router {
	app.AddRoutes(routes)
	router := app.GetRouter()
	return router
}

var routes = app.Routes{
	app.Route{
		"Index",
		"GET",
		"/",
		handler.Index,
	},
}
