package service

import (
	"context"

	"github.com/shiroemons/touhou_arrangement_chronicle/go/graph/model"
	"github.com/shiroemons/touhou_arrangement_chronicle/go/pkg/domain"
	"github.com/shiroemons/touhou_arrangement_chronicle/go/pkg/entity"
)

type EventSeriesService struct {
	esRepo domain.EventSeriesRepository
}

func EventSeriesServiceProvider(esRepo domain.EventSeriesRepository) *EventSeriesService {
	return &EventSeriesService{esRepo: esRepo}
}

func (s *EventSeriesService) New(ctx context.Context, input model.NewEventSeries) (*entity.EventSeries, error) {
	series := &entity.EventSeries{
		Name: input.Name,
	}
	if err := s.esRepo.Create(ctx, series); err != nil {
		return nil, err
	}
	return series, nil
}
