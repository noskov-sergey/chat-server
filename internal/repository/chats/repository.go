package chats

import "github.com/jackc/pgx/v5/pgxpool"

const (
	tableName = "chats"

	idColumn = "id"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewChatRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		db: db,
	}
}
