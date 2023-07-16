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

// create a new user using a username field nad return the user as a json object
func createUser(w http.ResponseWriter, r *http.Request) {
	user := User{
		Username: r.URL.Query().Get("username"),
	}

	query := `Insert Into users (username) Values (@username) returning id`
	args := pgx.NamedArgs{
		"username": user.Username,
	}

	err := db.QueryRow(context.Background(), query, args).Scan(&user.ID)
	handleError(err)

	json.NewEncoder(w).Encode(user)
}

func linkChatAndUser(w http.ResponseWriter, r *http.Request) {
	// get details of chat and user
	chat := Chat{
		ID: r.URL.Query().Get("chat"),
	}
	user := User{
		ID: r.URL.Query().Get("user"),
	}

	query := `insert into users_chat (chat_id, user_id) values (@chat_id, @user_id) `
	args := pgx.NamedArgs{
		"chat_id": chat.ID,
		"user_id": user.ID,
	}

	err := db.QueryRow(context.Background(), query, args).Scan()
	if err != pgx.ErrNoRows {
		http.ResponseWriter(w).WriteHeader(400)
	} else {
		http.ResponseWriter(w).WriteHeader(200)
	}
}

func handleRequests() {
	http.HandleFunc("/createChat", createChat)
	http.HandleFunc("/checkChat", checkChat)
	http.HandleFunc("/createUser", createUser)
	http.HandleFunc("/linkChatAndUser", linkChatAndUser)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
