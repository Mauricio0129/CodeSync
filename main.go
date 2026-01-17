package main

import (
	"CodeSync/routes"
	"context"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	connString := ""
	objcontext := context.Background()

	pool, err := pgxpool.New(objcontext, connString)
	if err != nil {
		log.Fatal("Error connecting to database", err)
	}
	defer pool.Close()

	routes.RegisterRoutes(pool)
	handler := routes.CorsMiddleware(http.DefaultServeMux)
	log.Fatal(http.ListenAndServe(":8080", handler))
}
