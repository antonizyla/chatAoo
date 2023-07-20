package user

import (
	"context"
	"errors"

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

func (r *Repository) createUser(name string) (usr User, err error) {
	var user User
	user.Name = name

	query := `Insert into users (name) values (@name) returning id`
	params := pgx.NamedArgs{
		"name": name,
	}

	er := r.DB.QueryRow(context.Background(), query, params).Scan(&user.ID)
	if er != nil {
		return user, er
	}
	if user.ID == uuid.Nil {
		return User{}, errors.New("user id is empty")
	}

	return user, nil
}

func (r *Repository) linkChatAndUser(chatId string, userId string) (err error) {
	query := `Insert into user_chat (chat_id, user_id) values (@chat_id, @user_id)`
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

func (r *Repository) getFromUsername(username string) (usr User, err error) {
	var user User

	user.Name = username

	query := `Select id from users where username = @username `
	params := pgx.NamedArgs{
		"username": username,
	}

	er := r.DB.QueryRow(context.Background(), query, params).Scan(&user.ID)
	if er != nil {
		return User{}, er
	}

	return user, nil
}

func (r *Repository) getFromUserID(userId string) (usr User, err error) {
	var user User

	user.ID, err = uuid.FromString(userId)
	if err != nil {
		return User{}, err
	}

	query := `Select username from users where id = @id `
	params := pgx.NamedArgs{
		"id": userId,
	}

	er := r.DB.QueryRow(context.Background(), query, params).Scan(&user.Name)
	if er != nil {
		return User{}, er
	}

	return user, nil
}
