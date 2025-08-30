package app

import (
	"AuthInGo/config/env"
	db "AuthInGo/db/repositories"
	"AuthInGo/router"
	"fmt"
	"net/http"
	"time"
)

// Config holds the configuration for the server
type Config struct {
	Addr string
}

// Application represents the application with its configuration
type Application struct {
	Config Config
	Store db.Storage
}

func NewConfig() Config {
	port := config.GetString("PORT", ":8080")
	return Config{
		Addr: port,
	}
}

func NewApplication(cfg Config) *Application {
	return &Application{
		Config: cfg,
		Store: *db.NewStorage(),
	}
}

func (app *Application) Run() error {
	server := &http.Server{
		Addr:         app.Config.Addr,
		Handler:      router.SetupRouter(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("Starting server on", app.Config.Addr)
	return server.ListenAndServe()
}