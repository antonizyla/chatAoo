package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid"
)

func (a *API) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Origin", "*")

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

func (a *API) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

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

func (a *API) UpdateUsername(w http.ResponseWriter, r *http.Request) {
	//set cors header to *
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Content-Type", "application/json")

    param := chi.URLParam(r, "id")

	var user User

	uid, err := uuid.FromString(param)
	if err != nil {
		http.Error(w, fmt.Sprintf("UUID formatted incorrectly %v", err.Error()), http.StatusBadRequest)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&user)
	user.ID = uid
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if user.Name == "" {
		http.Error(w, "name is empty", http.StatusBadRequest)
		return
	}

	_, err = a.repo.updateUserName(user.Name, uid)
	if err != nil {
		fmt.Println("error updating username")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.ResponseWriter(w).WriteHeader(http.StatusOK)

}

func (a *API) UsersChats(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "id")
	_, err := a.repo.getFromUserID(param)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	chats, err := a.repo.getLinkedChats(param)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// add cors header to *
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(chats)
}
