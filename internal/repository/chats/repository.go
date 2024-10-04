package chats

import (
	"github.com/noskov-sergey/platform-common/pkg/db"
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
