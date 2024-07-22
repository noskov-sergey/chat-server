package chats

import (
	"context"
	"fmt"
)

func (u *UseCase) Delete(ctx context.Context, chatID int) error {
	err := u.cRep.Delete(ctx, chatID)
	if err != nil {
		return fmt.Errorf("repository delete: %w", err)
	}

	return nil
}
