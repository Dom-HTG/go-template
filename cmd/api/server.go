package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Dom-HTG/go-template/internal/store"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Application struct {
	config config         // concrete implementation for application configurations.
	store  *store.Storage // concrete implementations for database interactions.
}

type config struct {
	addr string
}

func (app *Application) mount() *chi.Mux {
	router := chi.NewRouter()

	// base middleware stack.
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(time.Second * 60))

	// Register routes
	router.Get("/v1/healthcheck", app.HealthCheckHandler)

	// auth routes.
	router.Route("/v1/auth", func(router chi.Router) {
		router.Post("/email-signup", app.EmailPassSignupHandler)
		router.Post("/google-signup", app.GoogleOauthSignupHandler)
		router.Post("/facebook-signup", app.FacebookOauthSignupHandler)
	})
	// users routes.

	// posts routes.

	return router
}

func (app *Application) startServer(mux *chi.Mux) error {

	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 45, // timeout if server takes longer than 45s to write a response to the client.
		ReadTimeout:  time.Second * 15, // timeout if client takes longer than 15s to read the response from the server.
		IdleTimeout:  time.Minute,      // timeout after 60s when connection is idle.
	}

	log.Printf("Server is starting at http://localhost%s...", app.config.addr)

	return srv.ListenAndServe()
}
