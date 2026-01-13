package routes

import (
	"CodeSync/services"

	"github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router { // We create a router and add route handlers to it
	r := mux.NewRouter()
	r.HandleFunc("/test", services.H1)
	// Get all the user projects
	r.HandleFunc("/", services.H1).Methods("GET")
	// Create a new project
	r.HandleFunc("/", services.H1).Methods("POST")
	// Get entry level directory
	r.HandleFunc("/project/{project_id}", services.H1).Methods("GET")
	// New directory inside project (folder)
	r.HandleFunc("/project/{project_id}/directory", services.H1).Methods("POST")
	// Get directory (show nested directory content)
	r.HandleFunc("/project/{project_id}/directory/{directory_id}", services.H1).Methods("GET")
	// Create a new file inside a project
	r.HandleFunc("/project/{project_id}/file", services.H1).Methods("POST")
	// Get file to start editing it (upgrades to websocket protocol)
	r.HandleFunc("/ws/file/{fileId}", services.HandleWebSocket)
	return r
}
