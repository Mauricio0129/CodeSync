package services

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ProjectServices struct {
	Pool *pgxpool.Pool
}

func (s *ProjectServices) GetProjects(w http.ResponseWriter, r *http.Request) {

}
