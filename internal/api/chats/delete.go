package chats

import (
	"context"
	"fmt"
	desc "github.com/noskov-sergey/chat-server/pkg/chat_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	err := i.usecase.Delete(ctx, int(req.GetId()))
	if err != nil {
		return nil, fmt.Errorf("usecase delete: %w", err)
	}

	return &emptypb.Empty{}, nil
}
