package chat

import (
	"backend/api/resource/user"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid"
)

func (a *API) CreateChat(w http.ResponseWriter, r *http.Request) {
	// get chat from body
	chat := Chat{}
	err := json.NewDecoder(r.Body).Decode(&chat)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	chat, err = a.repo.Create(chat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(chat)
}

func (a *API) GetChat(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "chat_id")
	idU, err := uuid.FromString(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("The provided chat_id seems to malformed, make sure that it is a valid uuid: %v", err.Error()), http.StatusBadRequest)
		return
	}

	chat, err := a.repo.Get(idU)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error Fetching the chat from the database, make sure that the chat exists: %v", err.Error()), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(chat)
}

func (a *API) ChatsLinkedToUser(w http.ResponseWriter, r *http.Request) {
	user_id := chi.URLParam(r, "user_id")
	userUUID, err := uuid.FromString(user_id)
	if err != nil {
		http.Error(w, fmt.Sprintf("User Id Provided in the api call appears to be malformed or missing: %v", err.Error()), http.StatusBadRequest)
		return
	}
	chats, err := a.repo.GetChatsWithUser(userUUID)
	if err != nil {
		http.Error(w, fmt.Sprintf("An erorr occurred fetching the chats linked to the provided user: %v", err.Error()), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(chats)
}

func (a *API) DeleteChatLink(w http.ResponseWriter, r *http.Request) {
	chat_id := chi.URLParam(r, "chat_id")
	user_id := chi.URLParam(r, "user_id")

	chatUUID, err := uuid.FromString(chat_id)
	if err != nil {
		http.Error(w, fmt.Sprintf("The provided chat_id seems to be malformed, required to be a valid uuid: %v", err.Error()), http.StatusBadRequest)
		return
	}
	userUUID, err := uuid.FromString(user_id)
	if err != nil {
		http.Error(w, fmt.Sprintf("The provided user_id seems to be malformed, required to be a valid uuid: %v", err.Error()), http.StatusBadRequest)
		return
	}

	err = a.repo.DeleteChatLink(chatUUID, userUUID)
	if err != nil {
		http.Error(w, fmt.Sprintf("An error occurred when trying to delete the link between a user and their chats, make sure that the link actually exists: %v", err.Error()), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (a *API) Link(w http.ResponseWriter, r *http.Request) {

    fmt.Println("Linking user to chat")

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// get chat id and user id from request json body
	var link user.Link
	err := json.NewDecoder(r.Body).Decode(&link)
	if err != nil {
		http.Error(w, fmt.Sprintf("The provided body seems to be malformed, requires: {chat_id: 'valid uuid', user_id: 'valid uuid'}: %v", err.Error()), http.StatusBadRequest)
		return
	}

	err = a.repo.linkChatAndUser(link.ChatID, link.UserID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error occurred linking user and the chat: %v", err.Error()), http.StatusInternalServerError)
		return
	}
	http.ResponseWriter(w).WriteHeader(http.StatusOK)
}

func (a *API) GetUsersInChat(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// get chat id from Request
	chat_id := chi.URLParam(r, "chat_id")
	chatUUID, err := uuid.FromString(chat_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	users, err := a.repo.GetUsersInChat(chatUUID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(users)

}
