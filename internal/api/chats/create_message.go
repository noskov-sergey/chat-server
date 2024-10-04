package chats

import (
	"context"
	"fmt"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/noskov-sergey/chat-server/internal/converter"
	desc "github.com/noskov-sergey/chat-server/pkg/chat_v1"
)

func (i *Implementation) CreateMessage(ctx context.Context, req *desc.CreateMessageRequest) (*emptypb.Empty, error) {
	err := i.usecase.CreateMessage(ctx, converter.ToServiceFromMessages(req))
	if err != nil {
		return nil, fmt.Errorf("usecase create: %w", err)
	}

	return &emptypb.Empty{}, nil
}
