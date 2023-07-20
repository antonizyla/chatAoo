package router

import (
	"fmt"
	"net/http"

	"backend/api/resource/messages"
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
	router.Post("/users", usersApi.Create)
	router.Get("/users/{id}", usersApi.Get)
	router.Put("/users/link", usersApi.Link)

	// messages
	messageApi := message.New(db)
	router.Post("/messages", messageApi.Create)

	return router

}
