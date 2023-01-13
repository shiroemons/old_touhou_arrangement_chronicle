package entity

import (
	"time"

	"github.com/uptrace/bun"
)

type Track struct {
	bun.BaseModel `bun:"table:tracks,alias:t"`

	ID            string    `bun:",pk"`
	AlbumID       string    `bun:"album_id,nullzero,notnull"`
	Name          string    `bun:"name,nullzero,notnull"`
	DiscNumber    int       `bun:"disc_number,nullzero,notnull,default:1"`
	TrackNumber   int       `bun:"track_number,nullzero,notnull"`
	ReleaseDate   time.Time `bun:"release_date"`
	SearchEnabled bool      `bun:"search_enabled,nullzero,notnull,default:true"`
	CreatedAt     time.Time `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt     time.Time `bun:"updated_at,notnull,default:current_timestamp"`
}
