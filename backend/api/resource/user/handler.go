package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (a *API) Create(w http.ResponseWriter, r *http.Request) {
	var user User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if user.Name == "" {
		http.Error(w, "name is empty", http.StatusBadRequest)
		return
	}

	created, err := a.repo.createUser(user.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(created)
}

func (a *API) Link(w http.ResponseWriter, r *http.Request) {

	// get chat id and user id from request json body
	var link Link
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

func (a *API) Get(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "id")

	if len(param) == len(uuid.UUID{}.String()) {
		// get user by id
		user, err := a.repo.getFromUserID(param)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(user)
	} else {
		// get user by name
		user, err := a.repo.getFromUsername(param)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(user)
	}
}
