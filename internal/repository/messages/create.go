package messages

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/noskov-sergey/chat-server/internal/model"
	"github.com/noskov-sergey/platform-common/pkg/db"
)

func (r *Repository) Create(ctx context.Context, m model.Message) error {
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(userNameColumn, chatIdColumn, textColumn).
		Values(m.Username, m.ChatId, m.Text)

	sqlQuery, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("to sql: %w", err)
	}

	q := db.Query{
		Name:     "messages_repository.Create",
		QueryRaw: sqlQuery,
	}

	ct, err := r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return fmt.Errorf("exec: %w", err)
	}

	if ct.RowsAffected() == 0 {
		return fmt.Errorf("no rows are affected")
	}

	return nil
}
