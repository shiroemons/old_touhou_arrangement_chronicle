package domain

import (
	"context"

	"github.com/shiroemons/touhou_arrangement_chronicle/go/graph/model"
	"github.com/shiroemons/touhou_arrangement_chronicle/go/pkg/entity"
)

type ArtistService interface {
	New(ctx context.Context, input model.NewArtist) (*entity.Artist, error)
}

type ArtistRepository interface {
	Create(ctx context.Context, artist *entity.Artist) error
}
