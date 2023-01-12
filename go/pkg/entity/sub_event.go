package entity

import (
	"time"

	"github.com/uptrace/bun"
)

type SubEvent struct {
	bun.BaseModel `bun:"table:sub_events,alias:se"`

	ID        string    `bun:",pk"`
	EventID   string    `bun:"event_id,nullzero,notnull"`
	Name      string    `bun:"name,nullzero,notnull"`
	EventDate time.Time `bun:"event_date,nullzero,notnull"`
	CreatedAt time.Time `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:"updated_at,notnull,default:current_timestamp"`
}
