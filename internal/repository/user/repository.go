package user

import (
	"github.com/noskov-sergey/platform-common/pkg/db"
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
