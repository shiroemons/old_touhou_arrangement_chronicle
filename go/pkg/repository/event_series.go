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

func (r *EventSeriesRepository) Create(ctx context.Context, series *entity.EventSeries) error {
	if _, err := r.db.NewInsert().Model(series).Exec(ctx); err != nil {
		return err
	}
	return nil
}

func (r *EventSeriesRepository) GetMapInIDs(ctx context.Context, ids []string) (map[string]*entity.EventSeries, error) {
	eventSeries := make([]*entity.EventSeries, 0)
	err := r.db.NewSelect().Model(&eventSeries).Where("id IN (?)", bun.In(ids)).Scan(ctx)
	if err != nil {
		return nil, err
	}

	eventSeriesById := map[string]*entity.EventSeries{}
	for _, series := range eventSeries {
		s := series
		eventSeriesById[s.ID] = s
	}
	return eventSeriesById, nil
}
