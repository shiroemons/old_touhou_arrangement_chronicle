package repository

import (
	"context"

	"github.com/uptrace/bun"

	"github.com/shiroemons/touhou_arrangement_chronicle/go/pkg/entity"
)

type OriginalSongRepository struct {
	db *bun.DB
}

func OriginalSongRepositoryProvider(db *bun.DB) *OriginalSongRepository {
	return &OriginalSongRepository{db: db}
}

func (r *OriginalSongRepository) All(ctx context.Context) ([]*entity.OriginalSong, error) {
	originalSongs := make([]*entity.OriginalSong, 0)
	err := r.db.NewSelect().Model(&originalSongs).Order("id ASC").Scan(ctx)
	if err != nil {
		return nil, err
	}
	return originalSongs, nil
}
