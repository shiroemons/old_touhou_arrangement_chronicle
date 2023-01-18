package repository

import (
	"context"
	"database/sql"

	"github.com/shiroemons/touhou_arrangement_chronicle/go/pkg/ctxkey"
	"github.com/uptrace/bun"
)

type TxRepository struct {
	db *bun.DB
}

func TxRepositoryProvider(db *bun.DB) *TxRepository {
	return &TxRepository{db: db}
}

func (t *TxRepository) DoInTx(ctx context.Context, opts *sql.TxOptions, fn func(ctx context.Context, tx bun.Tx) error) error {
	tx, err := t.db.BeginTx(ctx, opts)
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
