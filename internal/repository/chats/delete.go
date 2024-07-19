package chats

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
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

	if _, err = r.db.Exec(ctx, sqlQuery, args...); err != nil {
		log.Printf("exec row: %v", err)
		return fmt.Errorf("exec row: %w", err)
	}

	return nil
}
