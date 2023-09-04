package message

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid"
	"github.com/olahol/melody"
)

func (a *API) DeleteMessage(w http.ResponseWriter, r *http.Request) {
	message_id := chi.URLParam(r, "message_id")
	msgUUID, err := uuid.FromString(message_id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Provided parameter seems to be malformed, %v", err.Error()), http.StatusBadRequest)
		return
	}
	_, err = a.repo.Delete(msgUUID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error deleting message, %v", err.Error()), http.StatusInternalServerError)
		return
	}
	http.ResponseWriter(w).WriteHeader(http.StatusOK)

}

func (a *API) Create(w http.ResponseWriter, r *http.Request) {
	// decode body into message
	// expects body, chat_id, user_id
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

	_, err = a.repo.Create(message)
	if err != nil {
		http.ResponseWriter(w).WriteHeader(http.StatusInternalServerError)
		http.ResponseWriter(w).Write([]byte("Error creating message"))
	}

	http.ResponseWriter(w).WriteHeader(http.StatusOK)

}

func (a *API) MessageWebsockets(s *melody.Session, msg []byte, lm *melody.Melody) {

	/*
				        Things that the user can do inside of a chat:
				            => editMessage | newMessage | deleteMessage | reaction

		                function expects one of the below

					   => {
					       actionType = editMessage,
					       message_id: uuid,
					       chat_id: uuid,
					       user_id: uuid,
					       new_body: string,
					   }

					   => {
					       actionType = newMessage,
					       chat_id: uuid,
					       user_id: uuid,
					       message_body: string
					   }

					   => {
					       actionType = deleteMessage
					       message_id: uuid
					   }

				       => {
				           actionType = reaction,
				           message_id = uuid,
				           user_id = uuid,
				           reaction_emoji = utf-8
				       }
	*/

	receivedMessage := WsRecievedMessage{}
	err := json.Unmarshal(msg, &receivedMessage)
	if err != nil {
		msg, _ = json.Marshal(`{"Status": "Error occurred unmarshalling message, 1"}`)
		goto broadcast
	}

	if receivedMessage.ActionType == "newMessage" {
		// check the values of the message
		if receivedMessage.ChatId == uuid.Nil || receivedMessage.MessageBody == "" || receivedMessage.UserId == uuid.Nil {
			msg, _ = json.Marshal(`{"Status": "Error occurred unmarshalling message, 2"}`)
			goto broadcast
		}
		// create the message in database and send the client the representation of the Message
		message := Message{
			Content: receivedMessage.MessageBody,
			ChatId:  receivedMessage.ChatId,
			UserId:  receivedMessage.UserId,
		}

		createdMessage, err := a.repo.Create(message)
		if err != nil {
			msg, _ = json.Marshal(`{"Status": "Error occurred creating message"}`)
		}
		msg, _ = json.Marshal(createdMessage)

	} else if receivedMessage.ActionType == "editMessage" {
		// todo
	} else if receivedMessage.ActionType == "deleteMessage" {
		fmt.Println("Deleting a message")

		// expects a message_id as uuid
		if receivedMessage.MessageId == uuid.Nil {
			msg, _ = json.Marshal(`{"Status": "message_id specified doesn't conform to uuid standard"}`)
			goto broadcast
		}

		deleted, err := a.repo.Delete(receivedMessage.MessageId)
		if err != nil {
			msg, _ = json.Marshal(`{"Status": "Error occurred deleting messsage from database"}`)
		}

		msg, _ = json.Marshal(deleted)

	} else if receivedMessage.ActionType == "addReaction" {
		fmt.Println("Reaction to a message")

		if receivedMessage.MessageId == uuid.Nil || receivedMessage.UserId == uuid.Nil {
			msg, _ = json.Marshal(`{"Status": "message_id or chat_id specified doesn't conform to uuid standard"}`)
			goto broadcast
		}

		err := a.repo.AddReaction(receivedMessage.MessageId, receivedMessage.UserId, receivedMessage.ReactionEmoji)
		if err != nil {
			msg, _ = json.Marshal(`{"Status": "could not create the reaction in the database"}`)
			goto broadcast
		}
		fmt.Printf("Sending New Reaction to client, %v \n", receivedMessage.ReactionEmoji)

	} else if receivedMessage.ActionType == "removeReaction" {

		fmt.Println("Removing a Reaction to a message")

		if receivedMessage.MessageId == uuid.Nil || receivedMessage.UserId == uuid.Nil {
			msg, _ = json.Marshal(`{"Status": "message_id or chat_id specified doesn't conform to uuid standard"}`)
			goto broadcast
		}

		err := a.repo.RemoveReaction(receivedMessage.MessageId, receivedMessage.UserId, receivedMessage.ReactionEmoji)
		if err != nil {
			msg, _ = json.Marshal(`{"Status": "could not delete the reaction in the database"}`)
			goto broadcast
		}

	} else {
		msg, _ = json.Marshal(`{"Status": "Error occurred unmarshalling message, type of action seems incorrect"}`)
	}

broadcast:
	lm.BroadcastFilter(msg, func(q *melody.Session) bool {
		return q.Request.URL.Query().Get("chat_id") == s.Request.URL.Query().Get("chat_id")
	})

}

func (a *API) GetReactions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	chat_id := chi.URLParam(r, "chat_id")
	chatUUID, err := uuid.FromString(chat_id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error occurred parsing time: %v", err.Error()), http.StatusBadRequest)
		return
	}
	reactions, err := a.repo.GetReactions(chatUUID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error occured fetching messages from database: %v", err.Error()), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(reactions)

}

func (a *API) GetMessages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// receives {chat_id}/{time_from}
	timefrom := chi.URLParam(r, "time_from")
	timefromAsInt, err := strconv.ParseInt(timefrom, 10, 64)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error occurred parsing time: %v", err.Error()), http.StatusBadRequest)
		return
	}
	time := time.UnixMilli(timefromAsInt)
	chatid := chi.URLParam(r, "chat_id")
	chatUUID, err := uuid.FromString(chatid)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error occured parsing chat: %v", err.Error()), http.StatusBadRequest)
		return
	}

	messages, err := a.repo.GetMessages(chatUUID, time)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error Occurred Fetching Messages from provided parameters: %v", err.Error()), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(messages)

}
