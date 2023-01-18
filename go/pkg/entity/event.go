package entity

import (
	"context"
	"time"

	"github.com/rs/xid"
	"github.com/uptrace/bun"

	"github.com/shiroemons/touhou_arrangement_chronicle/go/graph/model"
)

type Event struct {
	bun.BaseModel `bun:"table:events,alias:e"`

	ID            string    `bun:",pk"`
	Name          string    `bun:"name,nullzero,notnull"`
	EventSeriesID string    `bun:"event_series_id,nullzero,notnull"`
	CreatedAt     time.Time `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt     time.Time `bun:"updated_at,notnull,default:current_timestamp"`
}

var _ bun.BeforeAppendModelHook = (*Event)(nil)

func (e *Event) BeforeAppendModel(_ context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		if e.ID == "" {
			e.ID = xid.New().String()
		}
	}
	return nil
}

func (e *Event) ToGraphQL() *model.Event {
	return &model.Event{
		ID:   e.ID,
		Name: e.Name,
		EventSeries: &model.EventSeries{
			ID: e.EventSeriesID,
		},
	}
}
