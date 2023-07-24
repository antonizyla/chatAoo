package chat

import (
	"time"

	"github.com/gofrs/uuid"
)

type Chat struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

type User struct {
	UserID   uuid.UUID `json:"user_id"`
	UserName string    `json:"user_name"`
}
