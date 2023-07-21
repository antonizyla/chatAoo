package message

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/olahol/melody"
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

func (a *API) MessageWebsockets(s *melody.Session,  msg []byte, lm *melody.Melody) {

	var m Message
	err := json.Unmarshal(msg, &m)
	if err != nil {
		log.Fatalf("Unable to unmarshal message: %v", err)
	}

	m, err = a.repo.Create(&m)
	if err != nil {
		log.Fatalf("Error creating message: %v", err)
	}

	username, err := a.repo.userNameFromID(m.UserId)
	if err != nil {
		log.Fatalf("Error getting username: %v", err)
	}

	res := SentMessage{
		Message:  m,
		UserName: username,
	}

	stringRes, err := json.Marshal(res)
	if err != nil {
		log.Fatalf("Error marshalling message: %v", err)
	}

	lm.BroadcastFilter(stringRes, func(q *melody.Session) bool {
		return q.Request.URL.Query().Get("chat_id") == s.Request.URL.Query().Get("chat_id")
	})
}
