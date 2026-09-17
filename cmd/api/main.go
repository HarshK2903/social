package main

import (
	"log"
)

func main() {

	app, err := NewAPI()
	if err != nil {
		log.Fatal("failed to create application:", err)
	}

	func() {
		if err := app.Run(); err != nil {
			log.Fatal("server failed:", err)
		}
	}()

}
