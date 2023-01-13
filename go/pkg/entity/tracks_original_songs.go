package entity

import (
	"time"

	"github.com/uptrace/bun"
)

type TracksOriginalSongs struct {
	bun.BaseModel `bun:"table:tracks_original_songs,alias:tos"`

	TrackID        string    `bun:"track_id,pk,nullzero,notnull"`
	OriginalSongID string    `bun:"original_song_id,pk,nullzero,notnull"`
	CreatedAt      time.Time `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt      time.Time `bun:"updated_at,notnull,default:current_timestamp"`
}
