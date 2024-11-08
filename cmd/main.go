package main

import (
	"PIS_labs/internal/app"
	"PIS_labs/internal/interface/rest/handlers"
	"PIS_labs/internal/interface/rest/v1"
	"net/http"
)

func main() {
	controller := handlers.New()
	router := v1.MustGetGinRouter(controller)
	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}
	application := app.New(server)
	application.Start()
}
