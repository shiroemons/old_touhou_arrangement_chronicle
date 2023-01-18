package repository

import (
	"context"

	"github.com/uptrace/bun"

	"github.com/shiroemons/touhou_arrangement_chronicle/go/pkg/entity"
)

type EventSeriesRepository struct {
	db *bun.DB
}

func EventSeriesRepositoryProvider(db *bun.DB) *EventSeriesRepository {
	return &EventSeriesRepository{db: db}
}

func (es *EventSeriesRepository) Create(ctx context.Context, series *entity.EventSeries) error {
	if _, err := es.db.NewInsert().Model(series).Exec(ctx); err != nil {
		return err
	}
	return nil
}
