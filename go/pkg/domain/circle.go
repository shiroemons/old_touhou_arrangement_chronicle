package domain

import (
	"context"

	"github.com/shiroemons/touhou_arrangement_chronicle/go/graph/model"
	"github.com/shiroemons/touhou_arrangement_chronicle/go/pkg/entity"
)

type CircleService interface {
	New(ctx context.Context, input model.NewCircle) (*entity.Circle, error)
}

type CircleRepository interface {
	Create(ctx context.Context, circle *entity.Circle) error
}
