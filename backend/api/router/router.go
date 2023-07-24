package router

import (
	"fmt"
	"net/http"

	"backend/api/resource/chat"
	message "backend/api/resource/messages"
	"backend/api/resource/user"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/olahol/melody"
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
	router.Get("/users/{id}/linked-chats", usersApi.UsersChats)

	// messages
	m := melody.New()
	messageApi := message.New(db)
	//router.Post("/messages", messageApi.Create)
	router.Get("/messages/{chat_id}/{time_from}", messageApi.GetMessages)
	router.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		m.HandleRequest(w, r)
	})
	m.HandleMessage(func(s *melody.Session, b []byte) {
		messageApi.MessageWebsockets(s, b, m)
	})

	// chat
	chatApi := chat.New(db)
	router.Post("/chat", chatApi.CreateChat)
	router.Get("/chats/{chat_id}", chatApi.GetChat)
	router.Get("/chats/linked/{user_id}", chatApi.ChatsLinkedToUser)
	router.Delete("/chats/link/{chat_id}/{user_id}", chatApi.DeleteChatLink)
	router.Post("/chats/link", chatApi.Link)
	router.Get("/chats/{chat_id}/users", chatApi.GetUsersInChat)

	return router

}
