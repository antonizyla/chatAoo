package message

import (
	"github.com/gofrs/uuid"
	"time"
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
	UpdatedAt time.Time `json:"updated_at"`
	ChatId    uuid.UUID `json:"chat_id"`
	UserId    uuid.UUID `json:"user_id"`
	Deleted   bool      `json:"deleted"`
}

type SentMessage struct {
	Message  Message `json:"message"`
	UserName string  `json:"user_name"`
}

type DeletedMessage struct {
	MessageId uuid.UUID `json:"message_id"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ReactionToMessage struct {
	MessageID uuid.UUID `json:"message_id"`
	UserID    uuid.UUID `json:"user_id"`
	Reaction  string    `json:"reaction"`
}
