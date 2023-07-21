package user

import "github.com/gofrs/uuid"

type User struct {
	ID   uuid.UUID `json:"user_id"`
	Name string    `json:"name"`
}

type Link struct {
	ChatID uuid.UUID `json:"chat_id"`
	UserID uuid.UUID `json:"user_id"`
}
