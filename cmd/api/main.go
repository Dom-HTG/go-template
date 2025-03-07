package main

import "log"

func main() {
	// Setting up application configurations.

	// application configuration.
	cfg := config{
		addr: ":2719",
	}

	app := &application{
		config: cfg,
	}

	mux := app.mount()

	log.Fatal(app.startServer(mux))
}
