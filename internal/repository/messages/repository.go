package messages

import (
	"github.com/noskov-sergey/chat-server/internal/client/db"
)

const (
	tableName = "messages"

	chatIdColumn   = "chat_id"
	userNameColumn = "username"
	textColumn     = "text"
)

type Repository struct {
	db db.Client
}

func NewMessagesRepository(db db.Client) *Repository {
	return &Repository{
		db: db,
	}
}
