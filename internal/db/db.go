package db

import (
	"context"
	"database/sql"
)

type DBTX interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (result sql.Result, err error)
	PrepareContext(ctx context.Context, query string) (stmt *sql.Stmt, err error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (rows *sql.Rows, err error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) (row *sql.Row)
}

type Queries struct {
	db DBTX
}

func New(db DBTX) (queries *Queries) {
	return &Queries{db: db}
}

func (q *Queries) WithTx(tx *sql.Tx) (queries *Queries) {
	return &Queries{
		db: tx,
	}
}

func (q *Queries) ExecContext(ctx context.Context, query string, args ...interface{}) (result sql.Result, err error) {
	return q.db.ExecContext(ctx, query, args...)
}

func (q *Queries) PrepareContext(ctx context.Context, query string) (stmt *sql.Stmt, err error) {
	return q.db.PrepareContext(ctx, query)
}

func (q *Queries) QueryContext(ctx context.Context, query string, args ...interface{}) (rows *sql.Rows, err error) {
	return q.db.QueryContext(ctx, query, args...)
}

func (q *Queries) QueryRowContext(ctx context.Context, query string, args ...interface{}) (row *sql.Row) {
	return q.db.QueryRowContext(ctx, query, args...)
}
