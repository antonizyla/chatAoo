package chat

import "github.com/jackc/pgx/v5/pgxpool"

type API struct {
	repo *Repository
}

func New(db *pgxpool.Pool) *API {
	return &API{
		repo: NewRepo(db),
	}
}
