package main

import (
	"time"
)

type Chat struct {
	ID          string    `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	Created_At  time.Time `json:"created_at" db:"created_at"`
}

type User struct {
	ID       string `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
}

type UserChat struct {
	ChatID string `json:"chat_id" db:"chat_id"`
	UserID string `json:"user_id" db:"user_id"`
}
