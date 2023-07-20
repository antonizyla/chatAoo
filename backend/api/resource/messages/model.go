package message

import (
	"time"

	"github.com/gofrs/uuid"
)

type Message struct {
	ID        uuid.UUID `json:"message_id"`
	Content   string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	ChatId    uuid.UUID `json:"chat_id"`
	UserId    uuid.UUID `json:"user_id"`
}
