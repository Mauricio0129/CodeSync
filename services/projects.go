package services

import (
	"CodeSync/handlers"
	"CodeSync/schemas"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func rowToProject(row pgx.CollectableRow) (map[string]any, error) {
	var id, name string
	var createdAt, updatedAt time.Time

	err := row.Scan(&id, &name, &createdAt, &updatedAt)
	if err != nil {
		return nil, err
	}

	return map[string]any{
		"id":         id,
		"name":       name,
		"created_at": createdAt,
		"updated_at": updatedAt,
	}, nil
}

type ProjectServices struct {
	Pool *pgxpool.Pool
}

func (s *ProjectServices) GetProjects(w http.ResponseWriter, r *http.Request) {
	uncut_token := r.Header.Get("Authorization")
	cut_token := strings.TrimPrefix(uncut_token, "Bearer ")
	if cut_token == uncut_token || cut_token == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	context := r.Context()
	owner_id := handlers.VerifyToken(cut_token, w)
	results, err := s.Pool.Query(context, "SELECT id, name, created_at, updated_at  FROM projects WHERE owner_id = $1", owner_id)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	projects, err := pgx.CollectRows[map[string]any](results, rowToProject)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(projects)
}

func (s *ProjectServices) CreateProject(w http.ResponseWriter, r *http.Request) {
	data := schemas.RegisterProject{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Internal Server error", http.StatusInternalServerError)
		return
	}

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Authorization header required", http.StatusUnauthorized)
		return
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		http.Error(w, "Invalid authorization format", http.StatusUnauthorized)
		return
	}
	contxt := r.Context()
	token := parts[1]
	owner_id := handlers.VerifyToken(token, w)
	if owner_id == "" {
		return
	}

	result := s.Pool.QueryRow(contxt, "INSERT INTO projects (name, owner_id) VALUES ($1, $2) RETURNING id", data.Name, owner_id)

	internal_project_id := ""
	err = result.Scan(&internal_project_id)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	io.WriteString(w, "Project "+internal_project_id+" created")
}
