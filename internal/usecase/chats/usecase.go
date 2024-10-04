package chats

import (
	"context"

	"github.com/noskov-sergey/chat-server/internal/client/db"
	"github.com/noskov-sergey/chat-server/internal/model"
)

//go:generate mockgen -source usecase.go -destination mocks/mocks.go -typed true Repository
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

type UseCase struct {
	cRep      ChatRepository
	uRep      UserRepository
	mRep      MessageRepository
	txManager db.TxManager
}

func New(cRep ChatRepository, uRep UserRepository, mRep MessageRepository, txManager db.TxManager) *UseCase {
	return &UseCase{
		cRep:      cRep,
		uRep:      uRep,
		mRep:      mRep,
		txManager: txManager,
	}
}
