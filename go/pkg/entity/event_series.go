package entity

import (
	"time"

	"github.com/uptrace/bun"
)

type EventSeries struct {
	bun.BaseModel `bun:"table:event_series,alias:es"`

	ID        string    `bun:",pk"`
	Name      string    `bun:"name,nullzero,notnull"`
	CreatedAt time.Time `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:"updated_at,notnull,default:current_timestamp"`
}
