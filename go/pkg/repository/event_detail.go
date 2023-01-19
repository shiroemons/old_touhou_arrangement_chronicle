package repository

import (
	"context"

	"github.com/uptrace/bun"

	"github.com/shiroemons/touhou_arrangement_chronicle/go/pkg/ctxkey"
	"github.com/shiroemons/touhou_arrangement_chronicle/go/pkg/entity"
)

type EventDetailRepository struct {
	db *bun.DB
}

func EventDetailRepositoryProvider(db *bun.DB) *EventDetailRepository {
	return &EventDetailRepository{db: db}
}

func (r *EventDetailRepository) Create(ctx context.Context, eventDetail *entity.EventDetail) error {
	tx, ok := ctx.Value(ctxkey.TxCtxKey).(*bun.Tx)
	if ok {
		if _, err := tx.NewInsert().Model(eventDetail).Exec(ctx); err != nil {
			return err
		}
		return nil
	}
	if _, err := r.db.NewInsert().Model(eventDetail).Exec(ctx); err != nil {
		return err
	}
	return nil
}
