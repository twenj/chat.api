package main

import (
	"go-web/app"
	"chat.api/router"
	"log"
	"net/http"
)

func main() {
	routerInfo := start()
	log.Fatal(http.ListenAndServe(":8080", routerInfo))
}

func start() *app.AppRouter {
	router := router.Init()
	return router
}
