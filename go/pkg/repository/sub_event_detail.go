package repository

import (
	"context"

	"github.com/uptrace/bun"

	"github.com/shiroemons/touhou_arrangement_chronicle/go/pkg/ctxkey"
	"github.com/shiroemons/touhou_arrangement_chronicle/go/pkg/entity"
)

type SubEventDetailRepository struct {
	db *bun.DB
}

func SubEventDetailRepositoryProvider(db *bun.DB) *SubEventDetailRepository {
	return &SubEventDetailRepository{db: db}
}

func (r *SubEventDetailRepository) Create(ctx context.Context, detail *entity.SubEventDetail) error {
	tx, ok := ctx.Value(ctxkey.TxCtxKey).(*bun.Tx)
	if ok {
		if _, err := tx.NewInsert().Model(detail).Exec(ctx); err != nil {
			return err
		}
		return nil
	}
	if _, err := r.db.NewInsert().Model(detail).Exec(ctx); err != nil {
		return err
	}
	return nil
}
