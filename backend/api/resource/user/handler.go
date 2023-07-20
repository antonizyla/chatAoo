package user

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (a *API) Create(w http.ResponseWriter, r *http.Request) {
	created, err := a.repo.createUser(r.FormValue("name"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(created)
}

func (a *API) Link(w http.ResponseWriter, r *http.Request) {
	err := a.repo.linkChatAndUser(r.FormValue("chat_id"), r.FormValue("user_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.ResponseWriter(w).WriteHeader(http.StatusOK)
	http.ResponseWriter(w).Write([]byte("OK"))
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
