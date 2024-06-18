package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/Wasomeno/golang-practice/tutorial"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"

	"github.com/jackc/pgx/v5"
)

type queries struct {
	*tutorial.Queries
}

func main() {
	ctx := context.Background()

	godotenv.Load()

	portString := os.Getenv("PORT")

	if portString == "" {
		log.Fatal("PORT environment variable not set")
	}

	if portString == "" {
		log.Fatal("PORT environment variable not set")
	}

	connection, connectionError := pgx.Connect(ctx, "user=kevinananda dbname=postgres sslmode=disable")
	if connectionError != nil {
		log.Fatal(connectionError)
	}

	defer connection.Close(ctx)

	queries := queries{
		tutorial.New(connection),
	}

	router := chi.NewRouter()

	server := &http.Server{
		Addr:    ":" + portString,
		Handler: router,
	}

	router.Get("/ready", readinessHandler)
	router.Get("/error", errorHandler)
	router.Post("/users", queries.handleCreateUser)
	router.Get("/user", queries.handleGetUser)

	log.Printf("Listening on port %s", portString)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
