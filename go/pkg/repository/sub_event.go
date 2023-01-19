package repository

import (
	"context"

	"github.com/uptrace/bun"

	"github.com/shiroemons/touhou_arrangement_chronicle/go/pkg/ctxkey"
	"github.com/shiroemons/touhou_arrangement_chronicle/go/pkg/entity"
)

type SubEventRepository struct {
	db *bun.DB
}

func SubEventRepositoryProvider(db *bun.DB) *SubEventRepository {
	return &SubEventRepository{db: db}
}

func (se *SubEventRepository) Create(ctx context.Context, subEvent *entity.SubEvent) error {
	tx, ok := ctx.Value(ctxkey.TxCtxKey).(*bun.Tx)
	if ok {
		if _, err := tx.NewInsert().Model(subEvent).Exec(ctx); err != nil {
			return err
		}
		return nil
	}

	if _, err := se.db.NewInsert().Model(subEvent).Exec(ctx); err != nil {
		return err
	}
	return nil
}
