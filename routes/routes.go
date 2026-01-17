package routes

import (
	"CodeSync/services"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

func later(writer http.ResponseWriter, request *http.Request) {

}

// Simplest CORS - just set headers on everything
func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func RegisterRoutes(pool *pgxpool.Pool) {
	projects := &services.ProjectServices{Pool: pool}
	authentication := &services.AuthServices{Pool: pool}

	// Authentication
	http.HandleFunc("POST /register", authentication.RegisterUser)
	http.HandleFunc("POST /login", authentication.LoginUser)

	// Projects
	http.HandleFunc("GET /", projects.GetProjects)
	http.HandleFunc("POST /", projects.CreateProject)

	// Other routes...
	http.HandleFunc("GET /test", later)
}
