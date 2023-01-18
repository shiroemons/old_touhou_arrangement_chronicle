package domain

import (
	"context"

	"github.com/shiroemons/touhou_arrangement_chronicle/go/pkg/entity"
)

type EventDetailRepository interface {
	Create(ctx context.Context, ed *entity.EventDetail) error
}
