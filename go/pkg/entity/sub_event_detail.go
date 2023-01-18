package entity

import (
	"time"

	"github.com/uptrace/bun"
)

type SubEventDetail struct {
	bun.BaseModel `bun:"table:sub_event_details,alias:sed"`

	SubEventID  string    `bun:"sub_event_id,pk,nullzero,notnull"`
	EventDate   time.Time `bun:"event_date,nullzero"`
	EventStatus string    `bun:"event_status,nullzero,notnull,default:'scheduled'::event_status"`
	Description string    `bun:"description,nullzero,notnull,default:''"`
	CreatedAt   time.Time `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt   time.Time `bun:"updated_at,notnull,default:current_timestamp"`
}
