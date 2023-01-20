package service

import (
	"context"
	"database/sql"

	"github.com/uptrace/bun"

	"github.com/shiroemons/touhou_arrangement_chronicle/go/graph/model"
	"github.com/shiroemons/touhou_arrangement_chronicle/go/pkg/domain"
	"github.com/shiroemons/touhou_arrangement_chronicle/go/pkg/entity"
)

type CircleService struct {
	txRepo domain.TxRepository
	cRepo  domain.CircleRepository
	cdRepo domain.CircleDetailRepository
}

func CircleServiceProvider(txRepo domain.TxRepository, cRepo domain.CircleRepository, cdRepo domain.CircleDetailRepository) *CircleService {
	return &CircleService{txRepo: txRepo, cRepo: cRepo, cdRepo: cdRepo}
}

func (s *CircleService) New(ctx context.Context, input model.NewCircle) (*entity.Circle, error) {
	iType, iDetail := domain.InitialLetter(input.Name)
	circle := &entity.Circle{
		Name:                input.Name,
		InitialLetterType:   string(iType),
		InitialLetterDetail: iDetail,
	}
	err := s.txRepo.DoInTx(ctx, &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) error {
		if err := s.cRepo.Create(ctx, circle); err != nil {
			return err
		}

		detail := &entity.CircleDetail{CircleID: circle.ID}
		if err := s.cdRepo.Create(ctx, detail); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return circle, nil
}
