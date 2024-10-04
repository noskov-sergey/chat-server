package chats

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/noskov-sergey/platform-common/pkg/db"
	"log"
)

func (r *Repository) Delete(ctx context.Context, chatID int) error {
	builder := sq.
		Delete(tableName).
		PlaceholderFormat(sq.Dollar).
		Where("id = ?", chatID)

	sqlQuery, args, err := builder.ToSql()
	if err != nil {
		log.Fatalf("to sql: %v", err)
	}

	q := db.Query{
		Name:     "chats_repository.Delete",
		QueryRaw: sqlQuery,
	}

	if _, err = r.db.DB().ExecContext(ctx, q, args...); err != nil {
		return fmt.Errorf("exec row: %w", err)
	}

	return nil
}
