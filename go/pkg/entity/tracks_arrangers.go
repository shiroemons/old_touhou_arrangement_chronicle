package entity

import (
	"time"

	"github.com/uptrace/bun"
)

type TracksArrangers struct {
	bun.BaseModel `bun:"table:tracks_arrangers,alias:ta"`

	TrackID   string    `bun:"track_id,pk,nullzero,notnull"`
	ArtistID  string    `bun:"artist_id,pk,nullzero,notnull"`
	CreatedAt time.Time `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:"updated_at,notnull,default:current_timestamp"`
}
