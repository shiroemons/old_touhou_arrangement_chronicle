package entity

import (
	"context"
	"time"

	"github.com/rs/xid"
	"github.com/uptrace/bun"

	"github.com/shiroemons/touhou_arrangement_chronicle/go/graph/model"
)

type EventSeries struct {
	bun.BaseModel `bun:"table:event_series,alias:es"`

	ID        string    `bun:",pk"`
	Name      string    `bun:"name,nullzero,notnull"`
	CreatedAt time.Time `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:"updated_at,notnull,default:current_timestamp"`
}

var _ bun.BeforeAppendModelHook = (*EventSeries)(nil)

func (e *EventSeries) BeforeAppendModel(_ context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		if e.ID == "" {
			e.ID = xid.New().String()
		}
	}
	return nil
}

func (e *EventSeries) ToGraphQL() *model.EventSeries {
	return &model.EventSeries{
		ID:   e.ID,
		Name: e.Name,
	}
}

type EventSeriesList []*EventSeries

func (arr EventSeriesList) ToGraphQLs() []*model.EventSeries {
	res := make([]*model.EventSeries, len(arr))
	for i, es := range arr {
		res[i] = es.ToGraphQL()
	}
	return res
}
