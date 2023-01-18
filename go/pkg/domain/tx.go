package domain

import (
	"context"
	"database/sql"

	"github.com/uptrace/bun"
)

type TxRepository interface {
	DoInTx(ctx context.Context, opts *sql.TxOptions, fn func(ctx context.Context, tx bun.Tx) error) error
}
