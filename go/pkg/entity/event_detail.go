package entity

import (
	"time"

	"github.com/uptrace/bun"
)

type EventDetail struct {
	bun.BaseModel `bun:"table:event_details,alias:ed"`

	EventID     string      `bun:"event_id,pk,nullzero,notnull"`
	EventDates  []time.Time `bun:"event_dates,nullzero"`
	EventStatus string      `bun:"event_status,nullzero,notnull,default:'scheduled'::event_status"`
	Format      string      `bun:"format,nullzero,notnull,default:'offline'::event_format"`
	RegionCode  string      `bun:"region_code,nullzero,notnull,default:'JP'"`
	Address     string      `bun:"address,nullzero,notnull,default:''"`
	Description string      `bun:"description,nullzero,notnull,default:''"`
	URL         string      `bun:"url,nullzero,notnull,default:''"`
	TwitterURL  string      `bun:"twitter_url,nullzero,notnull,default:''"`
	CreatedAt   time.Time   `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt   time.Time   `bun:"updated_at,notnull,default:current_timestamp"`
}
