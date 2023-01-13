package entity

import (
	"time"

	"github.com/uptrace/bun"
)

type TrackISRC struct {
	bun.BaseModel `bun:"table:track_isrcs,alias:ti"`

	ID        string    `bun:",pk"`
	TrackID   string    `bun:"track_id,nullzero,notnull"`
	ISRC      string    `bun:"isrc,nullzero,notnull"`
	CreatedAt time.Time `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:"updated_at,notnull,default:current_timestamp"`
}
