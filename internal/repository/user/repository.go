package user

import "github.com/jackc/pgx/v5/pgxpool"

const (
	tableName = "chat_user"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		db: db,
	}
}
