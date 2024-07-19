package chats

import (
	"context"
	"github.com/noskov-sergey/chat-server/internal/model"
)

type ChatRepository interface {
	Create(ctx context.Context) (int, error)
	Delete(ctx context.Context, chatID int) error
}

type UserRepository interface {
	Create(ctx context.Context, u model.UserChat) error
}

type MessageRepository interface {
	Create(ctx context.Context, m model.Message) error
}

type useCase struct {
	cRep ChatRepository
	uRep UserRepository
	mRep MessageRepository
}

func New(cRep ChatRepository, uRep UserRepository, mRep MessageRepository) *useCase {
	return &useCase{
		cRep: cRep,
		uRep: uRep,
		mRep: mRep,
	}
}
