package service

import (
	"context"
	"database/sql"

	"github.com/uptrace/bun"

	"github.com/shiroemons/touhou_arrangement_chronicle/go/graph/model"
	"github.com/shiroemons/touhou_arrangement_chronicle/go/pkg/domain"
	"github.com/shiroemons/touhou_arrangement_chronicle/go/pkg/entity"
)

type EventService struct {
	txRepo domain.TxRepository
	eRepo  domain.EventRepository
	edRepo domain.EventDetailRepository
}

func EventServiceProvider(txRepo domain.TxRepository, eRepo domain.EventRepository, edRepo domain.EventDetailRepository) *EventService {
	return &EventService{txRepo: txRepo, eRepo: eRepo, edRepo: edRepo}
}

func (srv *EventService) New(ctx context.Context, input model.NewEvent) (*entity.Event, error) {
	event := &entity.Event{
		Name:          input.Name,
		EventSeriesID: input.EventSeriesID,
	}
	err := srv.txRepo.DoInTx(ctx, &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) error {
		if err := srv.eRepo.Create(ctx, event); err != nil {
			return err
		}

		detail := &entity.EventDetail{EventID: event.ID}
		if err := srv.edRepo.Create(ctx, detail); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return event, nil
}
