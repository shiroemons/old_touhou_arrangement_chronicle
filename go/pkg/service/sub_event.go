package service

import (
	"context"
	"database/sql"

	"github.com/uptrace/bun"

	"github.com/shiroemons/touhou_arrangement_chronicle/go/graph/model"
	"github.com/shiroemons/touhou_arrangement_chronicle/go/pkg/domain"
	"github.com/shiroemons/touhou_arrangement_chronicle/go/pkg/entity"
)

type SubEventService struct {
	txRepo  domain.TxRepository
	seRepo  domain.SubEventRepository
	sedRepo domain.SubEventDetailRepository
}

func SubEventServiceProvider(txRepo domain.TxRepository, seRepo domain.SubEventRepository, sedRepo domain.SubEventDetailRepository) *SubEventService {
	return &SubEventService{txRepo: txRepo, seRepo: seRepo, sedRepo: sedRepo}
}

func (srv *SubEventService) New(ctx context.Context, input model.NewSubEvent) (*entity.SubEvent, error) {
	sub := &entity.SubEvent{
		EventID: input.EventID,
		Name:    input.Name,
	}
	err := srv.txRepo.DoInTx(ctx, &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) error {
		if err := srv.seRepo.Create(ctx, sub); err != nil {
			return err
		}

		detail := &entity.SubEventDetail{SubEventID: sub.ID}
		if err := srv.sedRepo.Create(ctx, detail); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return sub, nil
}
