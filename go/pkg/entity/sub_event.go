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

func (se *SubEvent) BeforeAppendModel(_ context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		if se.ID == "" {
			se.ID = xid.New().String()
		}
	}
	return nil
}

func (se *SubEvent) ToGraphQL() *model.SubEvent {
	return &model.SubEvent{
		ID:   se.ID,
		Name: se.Name,
		Event: &model.Event{
			ID: se.EventID,
		},
	}
}
