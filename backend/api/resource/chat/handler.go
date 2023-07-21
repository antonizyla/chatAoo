package chat

import (
	"backend/api/resource/user"
	"encoding/json"
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
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	chat, err := a.repo.Get(idU)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(chat)
}

func (a *API) ChatsLinkedToUser(w http.ResponseWriter, r *http.Request) {
	user_id := chi.URLParam(r, "user_id")
	userUUID, err := uuid.FromString(user_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	chats, err := a.repo.GetChatsWithUser(userUUID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	http.Header.Set(w.Header(), "Content-Type", "application/json")
	http.Header.Set(w.Header(), "Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(chats)
}

func (a *API) DeleteChatLink(w http.ResponseWriter, r *http.Request) {
	chat_id := chi.URLParam(r, "chat_id")
	user_id := chi.URLParam(r, "user_id")

	chatUUID, err := uuid.FromString(chat_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	userUUID, err := uuid.FromString(user_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = a.repo.DeleteChatLink(chatUUID, userUUID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
        return
	}
	w.WriteHeader(http.StatusOK)
}

func (a *API) Link(w http.ResponseWriter, r *http.Request) {

	// get chat id and user id from request json body
	var link user.Link
	err := json.NewDecoder(r.Body).Decode(&link)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = a.repo.linkChatAndUser(link.ChatID, link.UserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.ResponseWriter(w).WriteHeader(http.StatusOK)
}
