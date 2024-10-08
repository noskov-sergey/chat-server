package chats

import (
	"context"
	"fmt"

	"github.com/noskov-sergey/chat-server/internal/converter"
	desc "github.com/noskov-sergey/chat-server/pkg/chat_v1"
)

func (i *Implementation) CreateChat(ctx context.Context, req *desc.CreateChatRequest) (*desc.CreateChatResponse, error) {
	id, err := i.usecase.CreateChat(ctx, converter.ToServiceFromChats(req))
	if err != nil {
		return nil, fmt.Errorf("usecase create: %w", err)
	}

	return &desc.CreateChatResponse{
		Id: int64(id),
	}, nil
}
