package chats

import (
	"github.com/noskov-sergey/chat-server/internal/client/db"
)

const (
	tableName = "chats"

	idColumn = "id"
)

type Repository struct {
	db db.Client
}

func NewChatRepository(db db.Client) *Repository {
	return &Repository{
		db: db,
	}
}
