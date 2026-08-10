package main

import (
	"log"

	"github.com/HarshK2903/social/internal/env"
)

func main() {
	if err := env.Load(".env"); err != nil {
		log.Fatal(err)
	}
	con := config{
		addr: env.GetString("ADDR", ":8080"),
	}
	app := &application{
		config: con,
	}
	mux := app.mount()
	log.Fatal(app.run(mux))

}
