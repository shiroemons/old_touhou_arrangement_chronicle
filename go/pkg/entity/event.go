package entity

import (
	"time"

	"github.com/uptrace/bun"
)

type Event struct {
	bun.BaseModel `bun:"table:events,alias:e"`

	ID            string    `bun:",pk"`
	Name          string    `bun:"name,nullzero,notnull"`
	EventSeriesID string    `bun:"event_series_id,nullzero,notnull"`
	CreatedAt     time.Time `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt     time.Time `bun:"updated_at,notnull,default:current_timestamp"`
}
