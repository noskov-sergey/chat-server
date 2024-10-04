package chats

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/noskov-sergey/platform-common/pkg/db"
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

	q := db.Query{
		Name:     "chats_repository.Create",
		QueryRaw: sqlQuery,
	}

	var insertedID int
	err = r.db.DB().ScanOneContext(ctx, &insertedID, q, args...)
	if err != nil {
		return 0, err
	}

	return insertedID, nil
}
