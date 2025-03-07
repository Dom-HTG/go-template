package main

import (
	"log"

	"github.com/Dom-HTG/go-template/internal/env"
	"github.com/Dom-HTG/go-template/internal/store"
	"github.com/joho/godotenv"
)

func main() {
	// Setting up application configurations.

	// Load env.
	err := godotenv.Load("app.env") // load environment variables from root directory.
	if err != nil {
		log.Fatalf("Error loading env file(s): %v", err)
	}

	// application configuration.
	cfg := config{
		addr: env.GetStringEnv("APP_PORT", ":8080"),
	}
	// Init store
	str := store.NewStorage(nil)

	app := &Application{
		config: cfg,
		store:  str,
	}

	mux := app.mount() // the mount method registers all of the applications middlewares and routes.

	log.Fatal(app.startServer(mux))
}
