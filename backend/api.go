package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5"
)

// create a new chat using a name and description field
func createChat(w http.ResponseWriter, r *http.Request) {

	chat := Chat{
		Name:        r.URL.Query().Get("name"),
		Description: r.URL.Query().Get("description"),
	}

	query := `Insert Into chat (name, description) Values (@name, @description) returning id, created_at`
	args := pgx.NamedArgs{
		"name":        chat.Name,
		"description": chat.Description,
	}

	err := db.QueryRow(context.Background(), query, args).Scan(&chat.ID, &chat.Created_At)
	handleError(err)

	json.NewEncoder(w).Encode(chat)

}

// check if a chat exists using the uuid of a chat
func checkChat(w http.ResponseWriter, r *http.Request) {
	chat := Chat{
		ID: r.URL.Query().Get("id"),
	}

	query := `Select name, description, created_at From chat Where id = @id`
	args := pgx.NamedArgs{
		"id": chat.ID,
	}

	err := db.QueryRow(context.Background(), query, args).Scan(&chat.Name, &chat.Description, &chat.Created_At)
	handleError(err)

	json.NewEncoder(w).Encode(chat)

}

func createUser(w http.ResponseWriter, r *http.Request) {
	user := User{
		Username: r.URL.Query().Get("username"),
        Chat: r.URL.Query().Get("chat"),
	}

	query := `Insert Into users (username, attached_chat) Values (@username, @chat) returning id`
	args := pgx.NamedArgs{
		"username": user.Username,
        "chat": user.Chat,
	}

	err := db.QueryRow(context.Background(), query, args).Scan(&user.ID)
	handleError(err)

	json.NewEncoder(w).Encode(user)
}

func exampleApi(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Hello World")
}

func handleRequests() {
	http.HandleFunc("/createChat", createChat)
	http.HandleFunc("/exampleApi", exampleApi)
	http.HandleFunc("/checkChat", checkChat)
    http.HandleFunc("/createUser", createUser)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
