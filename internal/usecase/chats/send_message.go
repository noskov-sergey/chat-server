package chats

import (
	"context"
	"fmt"
	"github.com/noskov-sergey/chat-server/internal/model"
)

func (u *UseCase) SendMessage(ctx context.Context, m model.Message) error {
	err := u.mRep.Create(ctx, m)
	if err != nil {
		return fmt.Errorf("repository create: %w", err)
	}

	return nil
}
