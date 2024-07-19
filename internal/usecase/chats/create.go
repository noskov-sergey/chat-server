package chats

import (
	"context"
	"fmt"
	"github.com/noskov-sergey/chat-server/internal/model"
)

func (u *useCase) Create(ctx context.Context, users model.Users) (int, error) {
	id, err := u.cRep.Create(ctx)
	if err != nil {
		return 0, fmt.Errorf("repository create: %w", err)
	}

	for _, userName := range users.Usernames {
		err = u.uRep.Create(
			ctx,
			model.UserChat{
				ChatID:   id,
				Username: userName,
			},
		)

		if err != nil {
			return 0, fmt.Errorf("user repository create: %w", err)
		}
	}

	return id, nil
}
