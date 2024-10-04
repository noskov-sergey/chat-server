package chats

import (
	"context"

	desc "github.com/noskov-sergey/chat-server/pkg/chat_v1"

	"github.com/noskov-sergey/chat-server/internal/model"
)

//go:generate mockgen -source service.go -destination mocks/mocks.go -typed true Usecase

type Usecase interface {
	CreateChat(ctx context.Context, users model.Users) (int, error)
	Delete(ctx context.Context, chatID int) error
	CreateMessage(ctx context.Context, m model.Message) error
}

type Implementation struct {
	desc.UnimplementedChatV1Server
	usecase Usecase
}

func New(u Usecase) *Implementation {
	return &Implementation{
		desc.UnimplementedChatV1Server{},
		u,
	}
}
