package router

import (
	"fmt"
	"net/http"

	"backend/api/resource/user"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func New(db *pgxpool.Pool) *chi.Mux {
	router := chi.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})

	// users
	usersApi := user.New(db)
	router.Post("/users/{name}", usersApi.Create)
	router.Get("/users/{id}", usersApi.Get)

	return router

}
