package entity

import (
	"time"

	"github.com/uptrace/bun"
)

type OriginalSongDistributionServiceURL struct {
	bun.BaseModel `bun:"table:original_song_distribution_service_urls,alias:osdsu"`

	ID             string    `bun:",pk"`
	OriginalSongID string    `bun:"original_song_id,nullzero,notnull"`
	Service        string    `bun:"service,nullzero,notnull"`
	URL            string    `bun:"url,nullzero,notnull"`
	CreatedAt      time.Time `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt      time.Time `bun:"updated_at,notnull,default:current_timestamp"`
}
