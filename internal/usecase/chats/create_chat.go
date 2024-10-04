package chats

import (
	"context"

	"github.com/noskov-sergey/chat-server/internal/model"
)

func (u *UseCase) CreateChat(ctx context.Context, users model.Users) (int, error) {
	var id int
	err := u.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		id, errTx = u.cRep.Create(ctx)
		if errTx != nil {
			return errTx
		}

		for _, userName := range users.Usernames {
			errTx = u.uRep.Create(
				ctx,
				model.UserChat{
					ChatID:   id,
					Username: userName,
				},
			)
			if errTx != nil {
				return errTx
			}
		}
		return nil
	})

	if err != nil {
		return 0, err
	}

	return id, nil
}
