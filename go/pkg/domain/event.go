package domain

import (
	"context"

	"github.com/shiroemons/touhou_arrangement_chronicle/go/graph/model"
	"github.com/shiroemons/touhou_arrangement_chronicle/go/pkg/entity"
)

type EventService interface {
	New(ctx context.Context, input model.NewEvent) (*entity.Event, error)
}

type EventRepository interface {
	Create(ctx context.Context, event *entity.Event) error
}
