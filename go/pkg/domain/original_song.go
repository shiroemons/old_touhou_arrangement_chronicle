package domain

import (
	"context"

	"github.com/shiroemons/touhou_arrangement_chronicle/go/pkg/entity"
)

type OriginalSongService interface {
	GetAll(ctx context.Context) (entity.OriginalSongs, error)
}

type OriginalSongRepository interface {
	All(ctx context.Context) ([]*entity.OriginalSong, error)
}
