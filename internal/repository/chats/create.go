package chats

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
)

func (r *Repository) Create(ctx context.Context) (int, error) {
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(idColumn).
		Values(sq.Expr("DEFAULT")).
		Suffix("RETURNING id")

	sqlQuery, args, err := builder.ToSql()
	if err != nil {
		return 0, fmt.Errorf("to sql: %w", err)
	}

	var insertedID int

	if err = r.db.QueryRow(ctx, sqlQuery, args...).Scan(&insertedID); err != nil {
		return 0, fmt.Errorf("query row: %w", err)
	}

	return insertedID, nil
}
