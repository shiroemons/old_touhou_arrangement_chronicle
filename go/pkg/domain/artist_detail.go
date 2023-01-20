package domain

import (
	"context"

	"github.com/shiroemons/touhou_arrangement_chronicle/go/pkg/entity"
)

type ArtistDetailRepository interface {
	Create(ctx context.Context, d *entity.ArtistDetail) error
}
