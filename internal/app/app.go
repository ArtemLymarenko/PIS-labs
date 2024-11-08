package app

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Application struct {
	server *http.Server
}

func New(server *http.Server) *Application {
	return &Application{
		server: server,
	}
}

func (app *Application) Start() {
	idleConnsClosed := make(chan struct{})

	go app.gracefulShutDown(idleConnsClosed)

	log.Println(fmt.Sprintf("HTTP server addr: %s", app.server.Addr))
	if err := app.server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		log.Println("HTTP server ListenAndServe: ", err)
	}

	<-idleConnsClosed

	log.Println("HTTP server stopped")
}

func (app *Application) gracefulShutDown(idleConnsClosed chan struct{}) {
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt, syscall.SIGTERM)
	<-sigint

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := app.server.Shutdown(ctx); err != nil {
		log.Println("http-server server shutdown: ", err)
	}

	close(idleConnsClosed)
}
