package message

import (
	"context"
	"fmt"
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

func (r *Repository) Delete(messageID uuid.UUID) (msg DeletedMessage, err error) {
	query := `UPDATE messages SET deleted = true, updated_at = NOW() WHERE id = @messageID returning id, updated_at`
	params := pgx.NamedArgs{
		"messageID": messageID,
	}
	msg = DeletedMessage{}
	err = r.DB.QueryRow(context.Background(), query, params).Scan(&msg.MessageId, &msg.UpdatedAt)
	if err != nil {
		return DeletedMessage{}, err
	}
	return msg, nil
}

func (r *Repository) Create(m Message) (message Message, err error) {

	query := `Insert into messages (message_body, chat_id, user_id) values (@message_body, @chat_id, @user_id) RETURNING id, created_at, updated_at`
	params := pgx.NamedArgs{
		"message_body": m.Content,
		"chat_id":      m.ChatId,
		"user_id":      m.UserId,
	}

	er := r.DB.QueryRow(context.Background(), query, params).Scan(&m.ID, &m.CreatedAt, &m.UpdatedAt)
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

func (r *Repository) GetReactions(chatId uuid.UUID) (reactions []ReactionToMessage, err error) {
	query := `Select message_id, user_id, emoji from reactions where message_id in (select message_id from messages where chat_id = @chat_id) `
	params := pgx.NamedArgs{
		"chat_id": chatId,
	}
	rows, err := r.DB.Query(context.Background(), query, params)
	if err != nil {
		return []ReactionToMessage{}, err
	}

	reactions = []ReactionToMessage{}
	for rows.Next() {
		react := ReactionToMessage{}
		err = rows.Scan(&react.MessageID, &react.UserID, &react.Reaction)
		if err != nil {
			return []ReactionToMessage{}, err
		}
		reactions = append(reactions, react)
	}

	return reactions, nil

}

func (r *Repository) GetMessages(chatID uuid.UUID, tFrom time.Time) (messages []Message, err error) {

	query := `SELECT id, message_body, created_at, updated_at, chat_id, user_id, deleted FROM messages WHERE chat_id = @chat_id AND created_at < @tFrom ORDER BY created_at LIMIT 100`
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

func (r *Repository) AddReaction(messageID uuid.UUID, userID uuid.UUID, reaction string) (err error) {
	query := `insert into reactions (message_id, user_id, emoji) values (@message_id, @user_id, @emoji)`
	params := pgx.NamedArgs{
		"message_id": messageID,
		"user_id":    userID,
		"emoji":      reaction,
	}
	_, err = r.DB.Query(context.Background(), query, params)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) RemoveReaction(messageID uuid.UUID, userID uuid.UUID, reaction string) (err error) {
	query := `delete from reactions where message_id = @message_id and user_id = @user_id and emoji = @emoji`
	params := pgx.NamedArgs{
		"message_id": messageID,
		"user_id":    userID,
		"emoji":      reaction,
	}

	_, err = r.DB.Query(context.Background(), query, params)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
