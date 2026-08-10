package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	msg := map[string]interface{}{
		"Status": "active",
		"time":   time.Now(),
	}
	byteData, err := json.Marshal(msg)
	if err != nil {
		log.Fatalf("Error while marshling data %v", err)
	}
	w.Write(byteData)
}
