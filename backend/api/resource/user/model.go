package user

import "github.com/gofrs/uuid"

type User struct {
	ID   uuid.UUID `json:"user_id"`
	Name string    `json:"name"`
}
