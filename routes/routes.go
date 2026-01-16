package routes

import (
	"CodeSync/services"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

func later(writer http.ResponseWriter, request *http.Request) {

}

func RegisterRoutes(pool *pgxpool.Pool) {
	// Initialize structures for services
	projects := &services.ProjectServices{Pool: pool}
	authentication := &services.AuthServices{Pool: pool}

	// Authentication:
	http.HandleFunc("POST /register", authentication.RegisterUser)
	http.HandleFunc("POST /login", authentication.LoginUser)

	// Projects:
	http.HandleFunc("GET /", projects.GetProjects) // Get all the projects
	http.HandleFunc("POST /", later)               // Register a new project

	// Directories:
	http.HandleFunc("GET /project/{project_id}", later)                          // Get top level content of a project
	http.HandleFunc("POST /project/{project_id}/directory", later)               // Create New directory inside project
	http.HandleFunc("GET /project/{project_id}/directory/{directory_id}", later) // Get directory content

	// Files:
	http.HandleFunc("POST /project/{project_id}/file", later) // Create a new file inside a project

	// Websocket
	http.HandleFunc("GET /ws/file/{fileId}", later) // Websocket endpoint to start editing a file

	// Test:
	http.HandleFunc("GET /test", later)
}
