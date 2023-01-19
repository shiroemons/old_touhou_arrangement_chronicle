package repository

import (
	"context"

	"github.com/uptrace/bun"

	"github.com/shiroemons/touhou_arrangement_chronicle/go/pkg/ctxkey"
	"github.com/shiroemons/touhou_arrangement_chronicle/go/pkg/entity"
)

type EventRepository struct {
	db *bun.DB
}

func EventRepositoryProvider(db *bun.DB) *EventRepository {
	return &EventRepository{db: db}
}

func (e *EventRepository) Create(ctx context.Context, event *entity.Event) error {
	tx, ok := ctx.Value(ctxkey.TxCtxKey).(*bun.Tx)
	if ok {
		if _, err := tx.NewInsert().Model(event).Exec(ctx); err != nil {
			return err
		}
		return nil
	}

	if _, err := e.db.NewInsert().Model(event).Exec(ctx); err != nil {
		return err
	}
	return nil
}

func (e *EventRepository) GetMapInIDs(ctx context.Context, ids []string) (map[string]*entity.Event, error) {
	events := make([]*entity.Event, 0)
	err := e.db.NewSelect().Model(&events).Where("id IN (?)", bun.In(ids)).Scan(ctx)
	if err != nil {
		return nil, err
	}

	eventById := map[string]*entity.Event{}
	for _, event := range events {
		ev := event
		eventById[ev.ID] = ev
	}
	return eventById, nil
}
