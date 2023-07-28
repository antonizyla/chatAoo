package message

import (
	"context"
	"time"

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

func (r *Repository) Delete(messageID uuid.UUID) (err error) {
	query := `UPDATE messages SET deleted = true, updated_at = NOW() WHERE id = @messageID`
	params := pgx.NamedArgs{
		"messageID": messageID,
	}
	_, err = r.DB.Exec(context.Background(), query, params)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) Create(m Message) (message Message, err error) {

	query := `Insert into messages (message_body, chat_id, user_id) values (@message_body, @chat_id, @user_id) RETURNING id, created_at`
	params := pgx.NamedArgs{
		"message_body": m.Content,
		"chat_id":      m.ChatId,
		"user_id":      m.UserId,
	}

	er := r.DB.QueryRow(context.Background(), query, params).Scan(&m.ID, &m.CreatedAt)
	if er != nil {
		return Message{}, er
	}
	return m, nil
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

func (r *Repository) GetMessages(chatID uuid.UUID, tFrom time.Time) (messages []Message, err error) {

	query := `SELECT id, message_body, created_at, updated_at, chat_id, user_id, deleted FROM messages WHERE chat_id = @chat_id AND created_at < @tFrom LIMIT 100 `
	params := pgx.NamedArgs{
		"chat_id": chatID,
		"tFrom":   tFrom,
	}

	rows, err := r.DB.Query(context.Background(), query, params)
	if err != nil {
		return []Message{}, err
	}

	messages = []Message{}
	for rows.Next() {
		var m Message
		err = rows.Scan(&m.ID, &m.Content, &m.CreatedAt, &m.UpdatedAt, &m.ChatId, &m.UserId, &m.Deleted)
		if err != nil {
			return []Message{}, err
		}
		messages = append(messages, m)
	}

	return messages, nil

}
