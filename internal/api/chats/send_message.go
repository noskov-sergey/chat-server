package chats

import (
	"context"
	"fmt"
	"github.com/noskov-sergey/chat-server/internal/converter"
	desc "github.com/noskov-sergey/chat-server/pkg/chat_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) SendMessage(ctx context.Context, req *desc.SendMessageRequest) (*emptypb.Empty, error) {
	err := i.usecase.SendMessage(ctx, converter.ToServiceFromMessages(req))
	if err != nil {
		return nil, fmt.Errorf("usecase create: %w", err)
	}

	return &emptypb.Empty{}, nil
}
