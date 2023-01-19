package domain

import (
	"context"

	"github.com/shiroemons/touhou_arrangement_chronicle/go/pkg/entity"
)

type SubEventDetailRepository interface {
	Create(ctx context.Context, sed *entity.SubEventDetail) error
}
