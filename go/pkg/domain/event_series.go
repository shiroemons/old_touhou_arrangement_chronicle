package domain

import (
	"context"

	"github.com/shiroemons/touhou_arrangement_chronicle/go/graph/model"
	"github.com/shiroemons/touhou_arrangement_chronicle/go/pkg/entity"
)

type EventSeriesService interface {
	New(ctx context.Context, input model.NewEventSeries) (*entity.EventSeries, error)
}

type EventSeriesRepository interface {
	Create(ctx context.Context, series *entity.EventSeries) error
}
