package chat

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

func (r *Repository) Create(chat Chat) (c Chat, err error) {
	query := `INSERT INTO chat (name, description) VALUES (@name, @description) RETURNING id, created_at`
	params := pgx.NamedArgs{
		"name":        chat.Name,
		"description": chat.Description,
	}
	er := r.DB.QueryRow(context.Background(), query, params).Scan(&chat.Id, &chat.CreatedAt)
	if er != nil {
		return Chat{}, er
	}
	return chat, nil
}

func (r *Repository) Get(id uuid.UUID) (c Chat, err error) {

	query := `SELECT  name, description, created_at FROM chat WHERE id = @id limit 100`
	params := pgx.NamedArgs{
		"id": id,
	}

	chat := Chat{}
	chat.Id = id
	er := r.DB.QueryRow(context.Background(), query, params).Scan(&chat.Name, &chat.Description, &chat.CreatedAt)
	if er != nil || er == pgx.ErrNoRows {
		return Chat{}, er
	}
	return chat, nil
}

func (r *Repository) GetChatsWithUser(user_id uuid.UUID) (chats []Chat, err error) {
	query := `select id, description, name, created_at from chat where id in (select chat_id from users_chat where user_id = @user_id)`
	params := pgx.NamedArgs{
		"user_id": user_id,
	}
	rows, err := r.DB.Query(context.Background(), query, params)
	if err != nil {
		return nil, err
	}
	// go through each row and scan into a chat struct
	chts := []Chat{}
	for rows.Next() {
		var chat Chat
		err = rows.Scan(&chat.Id, &chat.Description, &chat.Name, &chat.CreatedAt)
		if err != nil {
			return nil, err
		}
		chts = append(chts, chat)
	}
	return chts, nil
}

func (r *Repository) DeleteChatLink(chat_id uuid.UUID, user_id uuid.UUID) (err error) {
	query := `DELETE FROM users_chat WHERE chat_id = @chat_id AND user_id = @user_id`
	params := pgx.NamedArgs{
		"chat_id": chat_id,
		"user_id": user_id,
	}
	_, err = r.DB.Exec(context.Background(), query, params)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) linkChatAndUser(chatId uuid.UUID, userId uuid.UUID) (err error) {

	query := `Insert into users_chat (chat_id, user_id) values (@chat_id, @user_id)`
	params := pgx.NamedArgs{
		"chat_id": chatId,
		"user_id": userId,
	}

	_, er := r.DB.Exec(context.Background(), query, params)
	if er != nil {
		return er
	}

	return nil
}

func (r *Repository) GetUsersInChat(chatId uuid.UUID) (usrs []User, err error) {

	query := `select id, username from users where id in (select user_id from users_chat where chat_id = @chat_id)`
	params := pgx.NamedArgs{
		"chat_id": chatId,
	}

	rows, err := r.DB.Query(context.Background(), query, params)
	if err != nil {
		return nil, err
	}

	users := []User{}
	for rows.Next() {
		var user User
		err = rows.Scan(&user.UserID, &user.UserName)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil

}
