package services

import (
	"CodeSync/handlers"
	"CodeSync/schemas"
	"encoding/json"
	"io"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AuthServices struct {
	Pool *pgxpool.Pool
}

func (s *AuthServices) RegisterUser(w http.ResponseWriter, r *http.Request) {
	data := schemas.Register{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if data.Username == "" || data.Email == "" || data.Password == "" {
		http.Error(w, "All fields are required", http.StatusBadRequest)
		return
	}
	context := r.Context()

	log := s.Pool.QueryRow(
		context,
		"SELECT password_hash FROM users WHERE email = $1 OR username = $2",
		data.Email, data.Username,
	)

	err = log.Scan(nil)
	if err == nil {
		http.Error(w, "Email or username already taken", http.StatusConflict)
		return
	}

	if err != pgx.ErrNoRows {
		http.Error(w, "Internal Server error", http.StatusInternalServerError)
		return
	}

	hashed_password := handlers.HashPassword(data.Password, w)

	tag, err := s.Pool.Exec(context, "INSERT INTO users (email, username,  password_hash) VALUES ($1, $2, $3) ",
		data.Email, data.Username, hashed_password)

	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	if tag.RowsAffected() != 1 {
		http.Error(w, "User not created", http.StatusInternalServerError)
		return
	}

	io.WriteString(w, "successfully registered new user")
}

func (s *AuthServices) LoginUser(w http.ResponseWriter, r *http.Request) {
	var password string
	data := schemas.Login{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if data.Password == "" || (data.Email == "" && data.Username == "") {
		http.Error(w, "Password and either email or username required", http.StatusBadRequest)
		return
	}

	context := r.Context()

	log := s.Pool.QueryRow(
		context,
		"SELECT password_hash FROM users WHERE email = $1 OR username = $2",
		data.Email, data.Username,
	)

	err = log.Scan(&password)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	if !handlers.VerifyPassword(data.Password, password) {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	jwt := handlers.GenerateToken(w, data.Username)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"token": jwt,
	})
}
