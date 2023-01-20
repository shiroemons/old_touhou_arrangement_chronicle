package repository

import (
	"context"

	"github.com/uptrace/bun"

	"github.com/shiroemons/touhou_arrangement_chronicle/go/pkg/ctxkey"
	"github.com/shiroemons/touhou_arrangement_chronicle/go/pkg/entity"
)

type CircleRepository struct {
	db *bun.DB
}

func CircleRepositoryProvider(db *bun.DB) *CircleRepository {
	return &CircleRepository{db: db}
}

func (r *CircleRepository) Create(ctx context.Context, circle *entity.Circle) error {
	tx, ok := ctx.Value(ctxkey.TxCtxKey).(*bun.Tx)
	if ok {
		if _, err := tx.NewInsert().Model(circle).Exec(ctx); err != nil {
			return err
		}
		return nil
	}

	if _, err := r.db.NewInsert().Model(circle).Exec(ctx); err != nil {
		return err
	}
	return nil
}
