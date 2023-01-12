package entity

import (
	"time"

	"github.com/uptrace/bun"
)

type AlbumDistributionServiceURL struct {
	bun.BaseModel `bun:"table:album_distribution_service_urls,alias:aldsu"`

	ID        string    `bun:",pk"`
	AlbumID   string    `bun:"album_id,nullzero,notnull"`
	Service   string    `bun:"service,nullzero,notnull"`
	URL       string    `bun:"url,nullzero,notnull"`
	CreatedAt time.Time `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:"updated_at,notnull,default:current_timestamp"`
}
