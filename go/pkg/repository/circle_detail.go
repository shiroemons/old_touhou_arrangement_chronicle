package repository

import (
	"context"

	"github.com/uptrace/bun"

	"github.com/shiroemons/touhou_arrangement_chronicle/go/pkg/ctxkey"
	"github.com/shiroemons/touhou_arrangement_chronicle/go/pkg/entity"
)

type CircleDetailRepository struct {
	db *bun.DB
}

func CircleDetailRepositoryProvider(db *bun.DB) *CircleDetailRepository {
	return &CircleDetailRepository{db: db}
}

func (r *CircleDetailRepository) Create(ctx context.Context, circleDetail *entity.CircleDetail) error {
	tx, ok := ctx.Value(ctxkey.TxCtxKey).(*bun.Tx)
	if ok {
		if _, err := tx.NewInsert().Model(circleDetail).Exec(ctx); err != nil {
			return err
		}
		return nil
	}
	if _, err := r.db.NewInsert().Model(circleDetail).Exec(ctx); err != nil {
		return err
	}
	return nil
}
