package user

import (
	"github.com/noskov-sergey/chat-server/internal/client/db"
)

const (
	tableName = "chat_user"
)

type Repository struct {
	db db.Client
}

func NewUserRepository(db db.Client) *Repository {
	return &Repository{
		db: db,
	}
}
