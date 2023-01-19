package service

import (
	"context"

	"github.com/shiroemons/touhou_arrangement_chronicle/go/pkg/domain"
	"github.com/shiroemons/touhou_arrangement_chronicle/go/pkg/entity"
)

type OriginalSongService struct {
	osRepo domain.OriginalSongRepository
}

func OriginalSongServiceProvider(osRepo domain.OriginalSongRepository) *OriginalSongService {
	return &OriginalSongService{osRepo: osRepo}
}

func (s *OriginalSongService) GetAll(ctx context.Context) (entity.OriginalSongs, error) {
	products, err := s.osRepo.All(ctx)
	if err != nil {
		return make([]*entity.OriginalSong, 0), SrvErr(ctx, err.Error())
	}
	return products, nil
}
