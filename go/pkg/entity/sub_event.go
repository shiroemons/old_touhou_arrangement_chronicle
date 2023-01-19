package entity

import (
	"context"
	"time"

	"github.com/rs/xid"
	"github.com/uptrace/bun"

	"github.com/shiroemons/touhou_arrangement_chronicle/go/graph/model"
)

type SubEvent struct {
	bun.BaseModel `bun:"table:sub_events,alias:se"`

	ID        string    `bun:",pk"`
	EventID   string    `bun:"event_id,nullzero,notnull"`
	Name      string    `bun:"name,nullzero,notnull"`
	CreatedAt time.Time `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:"updated_at,notnull,default:current_timestamp"`
}

var _ bun.BeforeAppendModelHook = (*SubEvent)(nil)

func (e *SubEvent) BeforeAppendModel(_ context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		if e.ID == "" {
			e.ID = xid.New().String()
		}
	}
	return nil
}

func (e *SubEvent) ToGraphQL() *model.SubEvent {
	return &model.SubEvent{
		ID:   e.ID,
		Name: e.Name,
		Event: &model.Event{
			ID: e.EventID,
		},
	}
}
