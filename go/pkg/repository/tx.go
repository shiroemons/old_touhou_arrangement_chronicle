package repository

import (
	"context"
	"database/sql"

	"github.com/uptrace/bun"

	"github.com/shiroemons/touhou_arrangement_chronicle/go/pkg/ctxkey"
)

type TxRepository struct {
	db *bun.DB
}

func TxRepositoryProvider(db *bun.DB) *TxRepository {
	return &TxRepository{db: db}
}

func (r *TxRepository) DoInTx(ctx context.Context, opts *sql.TxOptions, fn func(ctx context.Context, tx bun.Tx) error) error {
	tx, err := r.db.BeginTx(ctx, opts)
	if err != nil {
		return err
	}

	c := context.WithValue(ctx, ctxkey.TxCtxKey, tx)

	var done bool

	defer func() {
		if !done {
			_ = tx.Rollback()
		}
	}()

	if err := fn(c, tx); err != nil {
		return err
	}

	done = true
	return tx.Commit()
}
