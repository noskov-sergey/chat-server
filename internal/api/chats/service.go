package chats

import (
	"context"
	"github.com/noskov-sergey/chat-server/internal/model"
	desc "github.com/noskov-sergey/chat-server/pkg/chat_v1"
)

type Usecase interface {
	Create(ctx context.Context, users model.Users) (int, error)
	Delete(ctx context.Context, chatID int) error
	SendMessage(ctx context.Context, m model.Message) error
}

type implementation struct {
	desc.UnimplementedChatV1Server
	usecase Usecase
}

func New(u Usecase) *implementation {
	return &implementation{
		desc.UnimplementedChatV1Server{},
		u,
	}
}
