package message

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	DB *pgxpool.Pool
}

func NewRepo(db *pgxpool.Pool) *Repository {
	return &Repository{
		DB: db,
	}
}

func (r *Repository) Create(m *Message) (message Message, err error) {

	query := `INSERT INTO messages (content, chat_id, user_id ) VALUES (  @content, @chat_id, @user_id ) returning id`
	params := pgx.NamedArgs{
		"content": m.Content,
		"chat_id": m.ChatId,
		"user_id": m.ChatId,
	}

	er := r.DB.QueryRow(context.Background(), query, params).Scan(m.ID)
	if er != nil {
		return Message{}, er
	}
	return *m, nil
}

func (r *Repository) userNameFromID(id uuid.UUID) (name string, err error) {
	query := `SELECT name FROM users WHERE id = @id`
	params := pgx.NamedArgs{
		"id": id,
	}
	er := r.DB.QueryRow(context.Background(), query, params).Scan(&name)
	if er != nil {
		return "", err
	}
	return name, nil
}
