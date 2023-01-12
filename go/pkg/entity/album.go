package entity

import (
	"time"

	"github.com/uptrace/bun"
)

type Album struct {
	bun.BaseModel `bun:"table:albums,alias:al"`

	ID                string    `bun:",pk"`
	Name              string    `bun:"name,nullzero,notnull"`
	ReleaseCircleName string    `bun:"release_circle_name,nullzero,notnull,default:''"`
	ReleaseDate       time.Time `bun:"release_date"`
	EventID           string    `bun:"event_id,nullzero,notnull,default:''"`
	SubEventID        string    `bun:"sub_event_id,nullzero,notnull,default:''"`
	SearchEnabled     bool      `bun:"search_enabled,nullzero,notnull,default:true"`
	CreatedAt         time.Time `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt         time.Time `bun:"updated_at,notnull,default:current_timestamp"`
}
