package messages

import "github.com/jackc/pgx/v5/pgxpool"

const (
	tableName = "messages"

	chatIdColumn   = "chat_id"
	userNameColumn = "username"
	textColumn     = "text"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewMessagesRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		db: db,
	}
}
