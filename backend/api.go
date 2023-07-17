package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/olahol/melody"
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

func getMessages(w http.ResponseWriter, r *http.Request) {
	timestamp := r.URL.Query().Get("timestamp")
	chatID := r.URL.Query().Get("chat_id")

	query := `Select id, chat_id, sender_id, content, created_at From messages Where chat_id = @chat_id And created_at < to_timestamp(@timestamp) And chat_id IS NOT NULL LIMIT 100`
	args := pgx.NamedArgs{
		"chat_id":   chatID,
		"timestamp": timestamp,
		"max":       100,
	}

	res := make([]map[string]interface{}, 0, 0)
	rows, err := db.Query(context.Background(), query, args)
	handleError(err)

	for rows.Next() {
		tmp := Message{}
		err := rows.Scan(&tmp.ID, &tmp.ChatID, &tmp.SenderID, &tmp.Content, &tmp.Created_At)
		if err != pgx.ErrNoRows {
			handleError(err)
		}
        tmpMap := map[string]interface{}{
            "id": tmp.ID,
            "chat_id": tmp.ChatID,
            "sender_id": tmp.SenderID,
            "content": tmp.Content,
            "created_at": tmp.Created_At,
            "sender_name": getUsernameFromID(tmp.SenderID),
        }
        res = append(res, tmpMap)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(res)

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

// get the name of a user from their id
func getUsernameFromID(id string) string {
	query := `Select username From users Where id = @id`
	args := pgx.NamedArgs{
		"id": id,
	}

	var username string
	err := db.QueryRow(context.Background(), query, args).Scan(&username)
	if err != pgx.ErrNoRows {
		handleError(err)
	}

	return username
}

func handleRequests() {

	m := melody.New()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		m.HandleRequest(w, r)
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		// perform work on msg - add to db etc

		var ms Message
		err := json.Unmarshal(msg, &ms)
		handleError(err)

		query := `Insert Into messages (chat_id, sender_id, content) Values (@chat_id, @sender_id, @content) returning id, created_at`
		args := pgx.NamedArgs{
			"chat_id":   ms.ChatID,
			"sender_id": ms.SenderID,
			"content":   ms.Content,
		}

		err = db.QueryRow(context.Background(), query, args).Scan(&ms.ID, &ms.Created_At)
		if err != pgx.ErrNoRows {
			handleError(err)
		}

		/* broadcasted format will be
		   {
		       "chat_id": string,
		       "sender_id": string,
		       "content": string,
		       "created_at": string,
		       "sender_name": string,
		   }
		*/

		res := map[string]interface{}{
			"chat_id":     ms.ChatID,
			"sender_id":   ms.SenderID,
			"content":     ms.Content,
			"created_at":  ms.Created_At,
			"sender_name": getUsernameFromID(ms.SenderID),
		}

		stringRes, err := json.Marshal(res)

		m.BroadcastFilter(stringRes, func(q *melody.Session) bool {
			return q.Request.URL.Query().Get("chat_id") == s.Request.URL.Query().Get("chat_id")
		})
	})

	http.HandleFunc("/createChat", createChat)
	http.HandleFunc("/checkChat", checkChat)
	http.HandleFunc("/createUser", createUser)
	http.HandleFunc("/linkChatAndUser", linkChatAndUser)
	http.HandleFunc("/getMessages", getMessages)

	log.Fatal(http.ListenAndServe(":8081", nil))
}
