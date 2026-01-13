package main

import (
	"CodeSync/routes"
	"context"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func main() {
	router := routes.RegisterRoutes()
	log.Println("Server starting on :8080")

	// Creates a Config object for connecting to the db
	poolConfig, err := pgxpool.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalln("Unable to parse DATABASE_URL:", err)
	}

	// Uses the context object with no wait limit which is handy for setup
	DB, err = pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		log.Fatalln("Unable to create connection pool:", err)
	}
	defer DB.Close() //we defer the closing to the very end when turning off

	log.Println("Database connected!")
	log.Fatal(http.ListenAndServe(":8080", router))
}
