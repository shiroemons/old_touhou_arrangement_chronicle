package service

import (
	"context"
	"database/sql"

	"github.com/uptrace/bun"

	"github.com/shiroemons/touhou_arrangement_chronicle/go/graph/model"
	"github.com/shiroemons/touhou_arrangement_chronicle/go/pkg/domain"
	"github.com/shiroemons/touhou_arrangement_chronicle/go/pkg/entity"
)

type ArtistService struct {
	txRepo domain.TxRepository
	aRepo  domain.ArtistRepository
	adRepo domain.ArtistDetailRepository
}

func ArtistServiceProvider(txRepo domain.TxRepository, aRepo domain.ArtistRepository, adRepo domain.ArtistDetailRepository) *ArtistService {
	return &ArtistService{txRepo: txRepo, aRepo: aRepo, adRepo: adRepo}
}

func (s *ArtistService) New(ctx context.Context, input model.NewArtist) (*entity.Artist, error) {
	iType, iDetail := domain.InitialLetter(input.Name)
	artist := &entity.Artist{
		Name:                input.Name,
		InitialLetterType:   string(iType),
		InitialLetterDetail: iDetail,
	}
	err := s.txRepo.DoInTx(ctx, &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) error {
		if err := s.aRepo.Create(ctx, artist); err != nil {
			return err
		}

		detail := &entity.ArtistDetail{ArtistID: artist.ID}
		if err := s.adRepo.Create(ctx, detail); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return artist, nil
}
