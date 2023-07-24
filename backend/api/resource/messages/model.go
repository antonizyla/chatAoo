package message

import (
	"time"

	"github.com/gofrs/uuid"
)

type WsRecievedMessage struct {
	ActionType    string    `json:"action_type"`
	MessageId     uuid.UUID `json:"message_id"`
	ChatId        uuid.UUID `json:"chat_id"`
	UserId        uuid.UUID `json:"user_id"`
	ReactionEmoji string    `json:"reaction_emoji"`
	MessageBody   string    `json:"message"`
}

type Message struct {
	ID        uuid.UUID `json:"message_id"`
	Content   string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	ChatId    uuid.UUID `json:"chat_id"`
	UserId    uuid.UUID `json:"user_id"`
}

type SentMessage struct {
	Message  Message `json:"message"`
	UserName string  `json:"user_name"`
}
