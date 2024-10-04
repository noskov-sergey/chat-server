package messages

import (
	"github.com/noskov-sergey/platform-common/pkg/db"
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
