package main

import (
	"log"

	"github.com/Dom-HTG/go-template/docs/database"
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
		db: dbConfig{
			dsn:            env.GetStringEnv("DB_DSN", "postgres://user:adminuserpassword@localhost/go-template?sslmode=disable"),
			maxOpenConn:    env.GetIntEnv("DB_MAX_OPEN_CONNS", 10),
			maxIdleConn:    env.GetIntEnv("DB_MAX_IDLE_CONNS", 5),
			maxIdleTimeout: env.GetStringEnv("DB_MAX_LIFETIME", "15min"),
		},
	}

	// connect to database.
	db, e := database.NewConn(cfg.db.dsn, cfg.db.maxOpenConn, cfg.db.maxIdleConn, cfg.db.maxIdleTimeout)
	if e != nil {
		log.Fatalf("Error connecting to database: %v", e)
	}

	// Init store
	str := store.NewStorage(db)

	app := &Application{
		config: cfg,
		store:  str,
	}

	mux := app.mount() // the mount method registers all of the applications middlewares and routes.

	log.Fatal(app.startServer(mux))
}
