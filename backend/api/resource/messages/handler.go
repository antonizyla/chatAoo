package message

import (
	"encoding/json"
	"net/http"
)

func (a *API) Create(w http.ResponseWriter, r *http.Request) {
	// receives {chat_id}/{user_id}/{content}
	// returns status

	// decode body into message
	var message Message
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if message.Content == "" {
		http.ResponseWriter(w).WriteHeader(http.StatusBadRequest)
		http.ResponseWriter(w).Write([]byte("Content is empty"))
	}

	_, err = a.repo.Create(&message)
	if err != nil {
		http.ResponseWriter(w).WriteHeader(http.StatusInternalServerError)
		http.ResponseWriter(w).Write([]byte("Error creating message"))
	}
	http.ResponseWriter(w).WriteHeader(http.StatusOK)

}
