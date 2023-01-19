package domain

import (
	"context"

	"github.com/shiroemons/touhou_arrangement_chronicle/go/graph/model"
	"github.com/shiroemons/touhou_arrangement_chronicle/go/pkg/entity"
)

type SubEventService interface {
	New(ctx context.Context, input model.NewSubEvent) (*entity.SubEvent, error)
}

type SubEventRepository interface {
	Create(ctx context.Context, subEvent *entity.SubEvent) error
}
