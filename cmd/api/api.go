package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/HarshK2903/social/internal/health"
)

type API struct {
	server       *Server
	mongoClient  *mongo.Client
	router       *chi.Mux
	healthHandle *health.Handler
}

func NewAPI() (*API, error) {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Println("warning: .env file not found")
	}
	// fmt.Print("Done!")
	mongoURI := os.Getenv("MONGO_URI")

	if mongoURI == "" {
		return nil, fmt.Errorf("MONGO_URI is not set")
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	//Connect to MongoDB

	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	defer cancel()

	mongoClient, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI(mongoURI),
	)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to connect to MongoDB: %w",
			err,
		)
	}
	if err := mongoClient.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf(
			"failed to ping MongoDB: %w",
			err,
		)
	}
	log.Println("MongoDB connected")

	dbName := os.Getenv("MONGO_DATABASE")

	if dbName == "" {
		dbName = "social"
	}
	db := mongoClient.Database(dbName)

	// Create health service and handler
	healthService := health.NewService(
	// mongoClient,
	)
	healthHandler := health.NewHandler(
		healthService,
	)

	//routers

	router := NewRoutes(
		healthHandler,
	)

	// reate HTTP server

	server := NewServer(
		port,
		router,
	)
	app := &API{
		server:       server,
		mongoClient:  mongoClient,
		router:       router,
		healthHandle: healthHandler,
	}

	_ = db

	return app, nil
}

func (a *API) Run() error {
	return a.server.Run()
}
