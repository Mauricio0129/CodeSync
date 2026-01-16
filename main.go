package main

import (
	"CodeSync/routes"
	"context"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	objcontext := context.Background()

	pool, err := pgxpool.New(objcontext, os.Getenv("DATABASE_URL")) // Logs the error and exits the program
	if err != nil {
		log.Fatal("Error connecting to database", err)
	}

	routes.RegisterRoutes(pool)
	log.Fatal(http.ListenAndServe(":8080", nil))

	defer pool.Close()
}
